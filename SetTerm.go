package main

var xmlTermGetLicensed = `
                <get-config>
                        <target>
                                <candidate />
                        </target>
                        <default-operation>merge</default-operation>
                        <config>
                                <configuration>
                                        <firewall>
                                                <family>
                                                        <inet>
                                                                <filter>
                                                                        <name>10mile</name>
                                                                        <term>
                                                                                <name>licensed</name>
                                                                        </term>
                                                                </filter>
                                                        </inet>
                                                </family>
                                        </firewall>
                        </configuration>
                        </config>
                </get-config>
        </rpc>
]]>]]>`


var xmlTermSetLicensed = `
                <edit-config>
                        <target>
                                <candidate />
                        </target>
                        <default-operation>merge</default-operation>
                        <config>
    				<configuration junos:changed-seconds="1608304598" junos:changed-localtime="date-time">
            				<firewall>
                				<family>
                    					<inet>
                        					<filter>
                            						<name>10mile</name>
                            						<term>
                                						<name>licensed</name>
                                						<from>
                                    							<address>
                                        							<ip-name>ip-address</ip-name>
                                    							</address>
                                						</from>
                            						</term>
                        					</filter>
                    					</inet>
                				</family>
            				</firewall>
    			</configuration>
                        </config>
                </edit-config>
        </rpc>
]]>]]>
<?xml version="1.0" encoding="UTF-8"?>
<rpc xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="2">
  <commit/>
</rpc>
]]>]]>
<?xml version="1.0" encoding="UTF-8"?>
<rpc xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="3">
  <close-session/>
</rpc>
]]>]]>`

var xmlTermSetNotLicensed = `
                <edit-config>
                        <target>
                                <candidate />
                        </target>
                        <default-operation>merge</default-operation>
                        <config>
                                <configuration junos:changed-seconds="1608304598" junos:changed-localtime="date-time">
                                        <firewall>
                                                <family>
                                                        <inet>
                                                                <filter>
                                                                        <name>10mile</name>
                                                                        <term>
                                                                                <name>notlicensed</name>
                                                                                <from>
                                                                                        <address>
                                                                                                <ip-name>ip-address</ip-name>
                                                                                        </address>
                                                                                </from>
                                                                        </term>
                                                                </filter>
                                                        </inet>
                                                </family>
                                        </firewall>
                        </configuration>
                        </config>
                </edit-config>
        </rpc>
]]>]]>
<?xml version="1.0" encoding="UTF-8"?>
<rpc xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="2">
  <commit/>
</rpc>
]]>]]>
<?xml version="1.0" encoding="UTF-8"?>
<rpc xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="3">
  <close-session/>
</rpc>
]]>]]>`


/*
 <config>
                                        <configuration>
                                                <routing-instances>
                                                        <instance operation="delete">
                                                                <name>?</name>
                                                        </instance>
                                                </routing-instances>
                                        </configuration>
                                </config>

*/

var xmlTermDeleteLicensed = `
                <edit-config>
                        <target>
                                <candidate />
                        </target>
                        <default-operation>merge</default-operation>
                        <config>
                                <configuration junos:changed-seconds="1608304598" junos:changed-localtime="date-time">
                                        <firewall>
                                                <family>
                                                        <inet>
                                                                <filter>
                                                                        <name>10mile</name>
                                                                        <term>
                                                                                <name>licensed</name>
                                                                                <from>
                                                                                        <address operation="delete">
                                                                                                <ip-name>ip-address</ip-name>
                                                                                        </address>
                                                                                </from>
                                                                        </term>
                                                                </filter>
                                                        </inet>
                                                </family>
                                        </firewall>
                        </configuration>
                        </config>
                </edit-config>
        </rpc>
]]>]]>
<?xml version="1.0" encoding="UTF-8"?>
<rpc xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="2">
  <commit/>
</rpc>
]]>]]>
<?xml version="1.0" encoding="UTF-8"?>
<rpc xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="3">
  <close-session/>
</rpc>
]]>]]>`


var xmlTermDeleteNotLicensed = `
                <edit-config>
                        <target>
                                <candidate />
                        </target>
                        <default-operation>merge</default-operation>
                        <config>
                                <configuration junos:changed-seconds="1608304598" junos:changed-localtime="date-time">
                                        <firewall>
                                                <family>
                                                        <inet>
                                                                <filter>
                                                                        <name>10mile</name>
                                                                        <term>
                                                                                <name>notlicensed</name>
                                                                                <from>
                                                                                        <address operation="delete">
                                                                                                <ip-name>ip-address</ip-name>
                                                                                        </address>
                                                                                </from>
                                                                        </term>
                                                                </filter>
                                                        </inet>
                                                </family>
                                        </firewall>
                        </configuration>
                        </config>
                </edit-config>
        </rpc>
]]>]]>
<?xml version="1.0" encoding="UTF-8"?>
<rpc xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="2">
  <commit/>
</rpc>
]]>]]>
<?xml version="1.0" encoding="UTF-8"?>
<rpc xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="3">
  <close-session/>
</rpc>
]]>]]>`




/*
var xmlTermSet = `
                <edit-config>
                        <target>
                                <candidate />
                        </target>
                        <default-operation>merge</default-operation>
                        <config>
                                <configuration junos:changed-seconds="1608304598" junos:changed-localtime="2020-12-18 15:16:38 GMT">
                                        <firewall>
                                                <family>
                                                        <inet>
                                                                <filter>
                                                                        <name>10mile</name>
                                                                        <term>
                                                                                <name>licensed</name>
                                                                                <from>
                                                                                        <address>
                                                                                                <name>10.1.4.0/24</name>
                                                                                        </address>
                                                                                </from>
                                                                        </term>
                                                                </filter>
                                                        </inet>
                                                </family>
                                        </firewall>
                        </configuration>
                        </config>
                </edit-config>
        </rpc>
]]>]]>`
*/

func GetXmlTermSetLicensed() string {
        return xmlTermSetLicensed
}

func GetXmlTermSetNotLicensed() string {
        return xmlTermSetNotLicensed
}

func GetXmlTermDeleteLicensed() string {
        return xmlTermDeleteLicensed
}

func GetXmlTermDeleteNotLicensed() string {
        return xmlTermDeleteNotLicensed
}

func GetXmlTermGetLicensed() string {
        return xmlTermGetLicensed
}
