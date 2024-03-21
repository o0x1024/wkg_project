package passive

import (
	"fmt"
	"strings"

	"golang.org/x/exp/maps"

	"wkg-agent/internal/assetCollect/subfinder/pkg/subscraping"
	"wkg-agent/internal/assetCollect/subfinder/pkg/subscraping/sources/alienvault"
	"wkg-agent/internal/assetCollect/subfinder/pkg/subscraping/sources/anubis"
	"wkg-agent/internal/assetCollect/subfinder/pkg/subscraping/sources/bevigil"
	"wkg-agent/internal/assetCollect/subfinder/pkg/subscraping/sources/binaryedge"
	"wkg-agent/internal/assetCollect/subfinder/pkg/subscraping/sources/bufferover"
	"wkg-agent/internal/assetCollect/subfinder/pkg/subscraping/sources/c99"
	"wkg-agent/internal/assetCollect/subfinder/pkg/subscraping/sources/censys"
	"wkg-agent/internal/assetCollect/subfinder/pkg/subscraping/sources/certspotter"
	"wkg-agent/internal/assetCollect/subfinder/pkg/subscraping/sources/chaos"
	"wkg-agent/internal/assetCollect/subfinder/pkg/subscraping/sources/chinaz"
	"wkg-agent/internal/assetCollect/subfinder/pkg/subscraping/sources/commoncrawl"
	"wkg-agent/internal/assetCollect/subfinder/pkg/subscraping/sources/crtsh"
	"wkg-agent/internal/assetCollect/subfinder/pkg/subscraping/sources/digitorus"
	"wkg-agent/internal/assetCollect/subfinder/pkg/subscraping/sources/dnsdb"
	"wkg-agent/internal/assetCollect/subfinder/pkg/subscraping/sources/dnsdumpster"
	"wkg-agent/internal/assetCollect/subfinder/pkg/subscraping/sources/dnsrepo"
	"wkg-agent/internal/assetCollect/subfinder/pkg/subscraping/sources/fofa"
	"wkg-agent/internal/assetCollect/subfinder/pkg/subscraping/sources/fullhunt"
	"wkg-agent/internal/assetCollect/subfinder/pkg/subscraping/sources/github"
	"wkg-agent/internal/assetCollect/subfinder/pkg/subscraping/sources/hackertarget"
	"wkg-agent/internal/assetCollect/subfinder/pkg/subscraping/sources/hunter"
	"wkg-agent/internal/assetCollect/subfinder/pkg/subscraping/sources/intelx"
	"wkg-agent/internal/assetCollect/subfinder/pkg/subscraping/sources/leakix"
	"wkg-agent/internal/assetCollect/subfinder/pkg/subscraping/sources/passivetotal"
	"wkg-agent/internal/assetCollect/subfinder/pkg/subscraping/sources/quake"
	"wkg-agent/internal/assetCollect/subfinder/pkg/subscraping/sources/rapiddns"
	"wkg-agent/internal/assetCollect/subfinder/pkg/subscraping/sources/riddler"
	"wkg-agent/internal/assetCollect/subfinder/pkg/subscraping/sources/robtex"
	"wkg-agent/internal/assetCollect/subfinder/pkg/subscraping/sources/securitytrails"
	"wkg-agent/internal/assetCollect/subfinder/pkg/subscraping/sources/shodan"
	"wkg-agent/internal/assetCollect/subfinder/pkg/subscraping/sources/sitedossier"
	"wkg-agent/internal/assetCollect/subfinder/pkg/subscraping/sources/threatbook"
	"wkg-agent/internal/assetCollect/subfinder/pkg/subscraping/sources/virustotal"
	"wkg-agent/internal/assetCollect/subfinder/pkg/subscraping/sources/waybackarchive"
	"wkg-agent/internal/assetCollect/subfinder/pkg/subscraping/sources/whoisxmlapi"
	"wkg-agent/internal/assetCollect/subfinder/pkg/subscraping/sources/zoomeye"
	"wkg-agent/internal/assetCollect/subfinder/pkg/subscraping/sources/zoomeyeapi"

	"github.com/projectdiscovery/gologger"
)

var AllSources = [...]subscraping.Source{
	&alienvault.Source{},
	&anubis.Source{},
	&bevigil.Source{},
	&binaryedge.Source{},
	&bufferover.Source{},
	&c99.Source{},
	&censys.Source{},
	&certspotter.Source{},
	&chaos.Source{},
	&chinaz.Source{},
	&commoncrawl.Source{},
	&crtsh.Source{},
	&digitorus.Source{},
	&dnsdb.Source{},
	&dnsdumpster.Source{},
	&dnsrepo.Source{},
	&fofa.Source{},
	&fullhunt.Source{},
	&github.Source{},
	&hackertarget.Source{},
	&hunter.Source{},
	&intelx.Source{},
	&passivetotal.Source{},
	&quake.Source{},
	&rapiddns.Source{},
	&riddler.Source{},
	&robtex.Source{},
	&securitytrails.Source{},
	&shodan.Source{},
	&sitedossier.Source{},
	&threatbook.Source{},
	&virustotal.Source{},
	&waybackarchive.Source{},
	&whoisxmlapi.Source{},
	&zoomeye.Source{},
	&zoomeyeapi.Source{},
	&leakix.Source{},
	// &threatminer.Source{}, // failing  api
	// &reconcloud.Source{}, // failing due to cloudflare bot protection
}

var NameSourceMap = make(map[string]subscraping.Source, len(AllSources))

func init() {
	for _, currentSource := range AllSources {
		NameSourceMap[strings.ToLower(currentSource.Name())] = currentSource
	}
}

// Agent is a struct for running passive subdomain enumeration
// against a given host. It wraps subscraping package and provides
// a layer to build upon.
type Agent struct {
	sources []subscraping.Source
}

// New creates a new agent for passive subdomain discovery
func New(sourceNames, excludedSourceNames []string, useAllSources, useSourcesSupportingRecurse bool) *Agent {
	sources := make(map[string]subscraping.Source, len(AllSources))

	if useAllSources {
		maps.Copy(sources, NameSourceMap)
	} else {
		if len(sourceNames) > 0 {
			for _, source := range sourceNames {
				if NameSourceMap[source] == nil {
					gologger.Warning().Msgf("There is no source with the name: %s", source)
				} else {
					sources[source] = NameSourceMap[source]
				}
			}
		} else {
			for _, currentSource := range AllSources {
				if currentSource.IsDefault() {
					sources[currentSource.Name()] = currentSource
				}
			}
		}
	}

	if len(excludedSourceNames) > 0 {
		for _, sourceName := range excludedSourceNames {
			delete(sources, sourceName)
		}
	}

	if useSourcesSupportingRecurse {
		for sourceName, source := range sources {
			if !source.HasRecursiveSupport() {
				delete(sources, sourceName)
			}
		}
	}

	gologger.Debug().Msgf(fmt.Sprintf("Selected source(s) for this search: %s", strings.Join(maps.Keys(sources), ", ")))

	// Create the agent, insert the sources and remove the excluded sources
	agent := &Agent{sources: maps.Values(sources)}

	return agent
}
