package stated

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"git.fractalqb.de/fractalqb/sllm"

	"github.com/CmdrVasquess/stated/att"
	"github.com/CmdrVasquess/stated/events"
	"github.com/CmdrVasquess/watched"
)

//go:generate versioner -pkg stated -bno build_no VERSION version.go

type HandlerFunc func(*EDState, events.Event) att.Change

var evtHdlrs = make(map[string]HandlerFunc)

const (
	ChgGame att.Change = (1 << iota)
	ChgCommander
	ChgSystem
	ChgLocation
	ChgShip

	ChgEND
)

func SaveJSON(file string, data interface{}, logTmpl string) error {
	if logTmpl != "" {
		log.Infoa(logTmpl, file)
	}
	tmp := file + "~"
	wr, err := os.Create(tmp)
	if err != nil {
		return err
	}
	defer wr.Close()
	enc := json.NewEncoder(wr)
	enc.SetIndent("", "\t")
	if err = enc.Encode(data); err != nil {
		return err
	}
	wr.Close()
	return os.Rename(tmp, file)
}

func LoadJSON(file string, allowEmpty bool, into interface{}, logTmpl string) error {
	if !strings.HasSuffix(file, ".json") {
		file = file + ".json"
	}
	if logTmpl != "" {
		log.Infoa(logTmpl, file)
	}
	rd, err := os.Open(file)
	switch {
	case allowEmpty && os.IsNotExist(err):
		log.Warna("`file` not exists, skip loading", file)
		return nil
	case err != nil:
		return err
	}
	defer rd.Close()
	dec := json.NewDecoder(rd)
	return dec.Decode(into)
}

type Config struct {
	CmdrFile        func(fid, name string) string
	ShutdownLogsOut bool
}

type CmdrFile struct {
	Dir    string
	MkDirs bool
}

func (cf CmdrFile) Filename(fid, _ string) string {
	if cf.MkDirs {
		if _, err := os.Stat(cf.Dir); os.IsNotExist(err) {
			os.MkdirAll(cf.Dir, 0666)
		}
	}
	file := fmt.Sprintf("EDstate-%s.json", fid)
	return filepath.Join(cf.Dir, file)
}

type ChangeEvent struct {
	Change att.Change
	Event  events.Event
}

type EDState struct {
	Config        `json:"-"`
	StatEDversion struct{ Major, Minor, Patch int }

	EDVersion string `json:"-"`
	Beta      bool   `json:"-"`
	Language  string `json:"-"`
	L10n      struct {
		Lang   string
		Region string
	} `json:"-"`
	Cmdr  *Commander
	Loc   JSONLocation
	Ships map[int]*Ship
	Mats  Materials

	Notify []chan<- ChangeEvent `json:"-"`

	lock sync.RWMutex
}

func NewEDState(cfg *Config) *EDState {
	res := &EDState{
		Ships: make(map[int]*Ship),
	}
	if cfg != nil {
		res.Config = *cfg
	}
	res.StatEDversion.Major = Major
	res.StatEDversion.Minor = Minor
	res.StatEDversion.Patch = Patch
	return res
}

func (ed *EDState) Reset() {
	ed.Cmdr = nil
	ed.Loc.Location = nil
	ed.Ships = make(map[int]*Ship)
}

func (es *EDState) SetEDVersion(v string) {
	es.EDVersion = v
	es.Beta = strings.Index(strings.ToLower(v), "beta") >= 0
}

var langMap = map[string]string{
	"English": "en",
	"German":  "de",
	"French":  "fr",
}

func ParseEDLang(edlang string) (lang, region string) {
	split := strings.Split(edlang, "\\")
	if len(split) != 2 {
		log.Errora("cannot parse `language`", edlang)
		return "", ""
	}
	lang = langMap[split[0]]
	if lang == "" {
		log.Warna("unknown `language`", split[0])
		return "", ""
	}
	return lang, split[1]
}

func (es *EDState) SetLanguage(lang string) {
	es.Language = lang
	es.L10n.Lang, es.L10n.Region = ParseEDLang(lang)
}

func (ed *EDState) SwitchCommander(fid string, name string) error {
	if err := ed.Save(); err != nil {
		log.Errore(err)
	}
	ed.Reset()
	if fid == "" {
		return nil
	}
	if name == "" {
		err := sllm.Error("empty commander name for `FID`", fid)
		log.Errore(err)
		return err
	}
	if ed.CmdrFile != nil {
		err := LoadJSON(ed.CmdrFile(fid, name), true, ed, "load ED state from `file`")
		if err != nil {
			log.Errore(err)
			return err
		}
	}
	if ed.Cmdr == nil {
		ed.Cmdr = &Commander{
			FID:  fid,
			Name: att.String(name),
		}
	} else if ed.Cmdr.FID != fid {
		err := sllm.Error("load `file with FID` for `FID`", ed.Cmdr.FID, fid)
		log.Errore(err)
		return err
	} else {
		ed.Cmdr.Name.Set(name, 0)
	}
	return nil
}

func (ed *EDState) MustCommander(where string) *Commander {
	if ed.Cmdr == nil {
		panic(fmt.Errorf("no current commander in '%s'", where))
	}
	return ed.Cmdr
}

func (es *EDState) RdLocked(do func() error) error {
	es.lock.RLock()
	defer es.lock.RUnlock()
	return do()
}

func (es *EDState) WrLocked(do func() error) error {
	es.lock.Lock()
	defer es.lock.Unlock()
	return do()
}

func (ed *EDState) Save() error {
	if ed.Cmdr == nil || ed.CmdrFile == nil {
		return nil
	}
	file := ed.CmdrFile(ed.Cmdr.FID, ed.Cmdr.Name.Get())
	err := SaveJSON(file, ed, "save ED state to `file`")
	return err
}

func (ed *EDState) Load(file string) error {
	return LoadJSON(file, true, ed, "load state from `file`")
}

func (ed *EDState) FindShip(id int) *Ship {
	return ed.Ships[id]
}

func (ed *EDState) GetShip(id int) *Ship {
	res := ed.FindShip(id)
	if res == nil {
		res = new(Ship)
		ed.Ships[id] = res
	}
	return res
}

func (ed *EDState) OnJournalEvent(e watched.JounalEvent) (err error) {
	defer func() {
		if p := recover(); p != nil {
			switch x := p.(type) {
			case error:
				err = x
			case string:
				err = errors.New(x)
			default:
				err = fmt.Errorf("%+v", x)
			}
		}
	}()
	event, err := e.Event.PeekEvent()
	if err != nil {
		return err
	}
	etype := events.EventType(event)
	if etype == nil {
		log.Debuga("unknown `journal event`", event)
		return nil
	}
	eh := evtHdlrs[event]
	if eh == nil {
		log.Debuga("no handler for `journal event`", event)
		return nil
	}
	evt := etype.New()
	if err = json.Unmarshal(e.Event, evt); err != nil {
		return err
	}
	chg := eh(ed, evt)
	ed.ntfChg(chg, evt)
	return err
}

func (ed *EDState) OnStatusEvent(e watched.StatusEvent) error {
	etype := events.EventType(e.Type.String())
	if etype == nil {
		log.Debuga("unknown `status event`", e.Type.String())
		return nil
	}
	// TODO status event
	return errors.New("NYI: EDState status event")
}

func (ed *EDState) Close() error {
	return ed.Save()
}

func (ed *EDState) ntfChg(chg att.Change, e events.Event) {
	ce := ChangeEvent{chg, e}
	for i, c := range ed.Notify {
		select {
		case c <- ce:
			log.Tracea("sent `change` to `listener`", chg, i)
		default:
			log.Tracea("drop `change` for blocking `listener`", chg, i)
		}
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
