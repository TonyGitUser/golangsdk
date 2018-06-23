/*
Before adding backend ECSs, you need to check whether the rule of the security group to which the ECSs belong allows access by 100.125.0.0/16, and specify the protocol and port for health checks. If the protocol and port are not specified, the health of the added ECSs cannot be checked.

Sample Code, This interface is used to add backend ECSs to a listener for monitoring.


    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    listenerId := "d5f3197c6bd8491ca1dfc905350d85ea"
    result, err := backendces.Create(client, tenantId, listenerId, backendces.CreateOpts{
        Items: []struct {
            ServerId string `json:"server_id,omitempty"`
            Address  string `json:"address,omitempty"`
        }{
            {
                ServerId: "e1040c67-b130-41ee-9405-07c6c6c20208",
                Address:  "192.168.1.215",
            },
        },
    }).Extract()

    if err != nil {
        panic(err)
    }

Sample Code, This interface is used to remove backend ECSs from a listener. Multiple backend ECSs can be removed concurrently.


    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    listenerId := "ca082827b61d4902bfaf32e8174e244a"
    backendces.Delete(client, tenantId, listenerId, backendces.DeleteOpts{
        RemoveMember: []struct {
            Id string `json:"id,omitempty"`
        }{
            {
                Id: "e1040c67-b130-41ee-9405-07c6c6c20208",
            },
        },
    })

    if err != nil {
        panic(err)
    }

Sample Code, This interface is used to query backend ECSs of a listener. Common tenants can use this interface to list the backend ECSs, while administrator tenants can only see an empty list when using this interface to query backend ECSs.


    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    listenerId := "d5f3197c6bd8491ca1dfc905350d85ea"
    result, err := backendces.List(client, tenantId, listenerId, backendces.ListOpts{
        Limit: "1",
    }).Extract()

    if err != nil {
        panic(err)
    }
*/
package backendces
