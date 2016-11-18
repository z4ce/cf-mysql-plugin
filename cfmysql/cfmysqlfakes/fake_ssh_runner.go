// This file was generated by counterfeiter
package cfmysqlfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/plugin"
	"github.com/andreasf/cf-mysql-plugin/cfmysql"
)

type FakeSshRunner struct {
	OpenSshTunnelStub        func(cliConnection plugin.CliConnection, toService cfmysql.MysqlService, throughApp string, localPort int)
	openSshTunnelMutex       sync.RWMutex
	openSshTunnelArgsForCall []struct {
		cliConnection plugin.CliConnection
		toService     cfmysql.MysqlService
		throughApp    string
		localPort     int
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSshRunner) OpenSshTunnel(cliConnection plugin.CliConnection, toService cfmysql.MysqlService, throughApp string, localPort int) {
	fake.openSshTunnelMutex.Lock()
	fake.openSshTunnelArgsForCall = append(fake.openSshTunnelArgsForCall, struct {
		cliConnection plugin.CliConnection
		toService     cfmysql.MysqlService
		throughApp    string
		localPort     int
	}{cliConnection, toService, throughApp, localPort})
	fake.recordInvocation("OpenSshTunnel", []interface{}{cliConnection, toService, throughApp, localPort})
	fake.openSshTunnelMutex.Unlock()
	if fake.OpenSshTunnelStub != nil {
		fake.OpenSshTunnelStub(cliConnection, toService, throughApp, localPort)
	}
}

func (fake *FakeSshRunner) OpenSshTunnelCallCount() int {
	fake.openSshTunnelMutex.RLock()
	defer fake.openSshTunnelMutex.RUnlock()
	return len(fake.openSshTunnelArgsForCall)
}

func (fake *FakeSshRunner) OpenSshTunnelArgsForCall(i int) (plugin.CliConnection, cfmysql.MysqlService, string, int) {
	fake.openSshTunnelMutex.RLock()
	defer fake.openSshTunnelMutex.RUnlock()
	return fake.openSshTunnelArgsForCall[i].cliConnection, fake.openSshTunnelArgsForCall[i].toService, fake.openSshTunnelArgsForCall[i].throughApp, fake.openSshTunnelArgsForCall[i].localPort
}

func (fake *FakeSshRunner) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.openSshTunnelMutex.RLock()
	defer fake.openSshTunnelMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeSshRunner) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ cfmysql.SshRunner = new(FakeSshRunner)