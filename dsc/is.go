package dsc

import (
	"fmt"
	"strings"
	"sync"

	"github.com/nelsonsaake/go/tr"

	"github.com/nelsonsaake/go/is"

	"github.com/nelsonsaake/go/cfgs"
)

var descriptorService *DescriptorService

func getDefaultDescriptorService() *DescriptorService {
	if descriptorService == nil {
		descriptorService = NewDescriptorService()
	}
	return descriptorService
}

type DescriptorService struct {
	rulesCache    sync.Map // map[string][]map[string]any
	propertyCache sync.Map // map[string]bool
}

func NewDescriptorService() *DescriptorService {
	return &DescriptorService{}
}

func (ds *DescriptorService) getRules(rsrc ...string) []map[string]any {

	key := strings.Join(rsrc, "|")

	if cache, ok := ds.rulesCache.Load(key); ok {
		return cache.([]map[string]any)
	}

	// get the rules
	var rules []map[string]any
	cfgs.GetAs(&rules, rsrc...)

	// cache the rules for subsequent use
	ds.rulesCache.Store(key, rules)

	return rules
}

func (ds *DescriptorService) Is(obj is.Subject, id string, rsrc ...string) bool {

	ck := id + "|" + strings.Join(rsrc, "|")

	if cache, ok := ds.propertyCache.Load(ck); ok {
		return cache.(bool)
	}

	rules := ds.getRules(rsrc...)

	res := is.MatchesAny(obj, rules)
	ds.propertyCache.Store(ck, res)

	return res
}

func (ds *DescriptorService) Tr(obj tr.Subject, id, t string, rsrc ...string) any {

	ck := id + "|" + strings.Join(rsrc, "|")

	if cache, ok := ds.propertyCache.Load(ck); ok {
		return cache
	}

	rules := ds.getRules(rsrc...)
	if len(rules) == 0 {
		panic(fmt.Errorf("rules cannot be empty: %v", rsrc))
	}

	res := tr.MatchesAny(obj, t, rules)

	ds.propertyCache.Store(ck, res)

	return res
}

func Is(obj is.Subject, id string, rsrc ...string) bool {

	return getDefaultDescriptorService().Is(obj, id, rsrc...)
}
