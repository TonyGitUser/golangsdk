/*
A listener specifies the load balancer protocol and port, ECS protocol and port, and listening policy.

Sample Code, This interface is used to create a listener for a load balancer to monitor status of backend ECSs.


    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result, err := listeners.Create(client, tenantID, listeners.CreateOpts{
        Name:              "TestLS",
        Description:       "LS TESTING",
        LoadbalancerId:    "ebe982b8afe04851bbc26ad4609003bf",
        Protocol:          "HTTP",
        Port:              80,
        BackendProtocol:   "HTTP",
        BackendPort:       81,
        LbAlgorithm:       "roundrobin",
        SessionSticky:     true,
        StickySessionType: "insert",
        CookieTimeout:     1,
    }).Extract()

    if err != nil {
        panic(err)
    }

Sample Code, This interface is used to delete a listener.


    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    listenerId := "223a537e337e4f9b947b8b39eb1b1f6c"
    listeners.Delete(client, tenantID, listenerId)

Sample Code, This interface is used to query details about a listener.


    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    listenerId := "223a537e337e4f9b947b8b39eb1b1f6c"
    result, err := listeners.Get(client, tenantID, listenerId).Extract()

    if err != nil {
        panic(err)
    }

Sample Code, This interface is used to query listeners using search criteria and to display the listeners in a list.


    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result, err := listeners.List(client, tenantID, listeners.ListOpts{
       LoadbalancerId: "ebe982b8afe04851bbc26ad4609003bf",
    }).Extract()

Sample Code, This interface is used to modify the name, description, and management status of a listener.


    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    listenerId := "223a537e337e4f9b947b8b39eb1b1f6c"
    result, err := listeners.Update(client, tenantID, listenerId, listeners.UpdateOpts{
        Name:        "ModifiedTestLS",
        Description: "LS TESTING",
        Port:        80,
        BackendPort: 81,
        LbAlgorithm: "roundrobin",
    }).Extract()

    if err != nil {
        panic(err)
    }
*/
package listeners
