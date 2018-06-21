/*
This interface is used to manage and perform operations on network resources, including querying networks, creating a network, querying a specified network, deleting a network, and updating a network.

Sample Code, Creates a network.


    result, err := networks.Create(client, networks.CreateOpts{
        Network: networks.CreateNetwork{
            Name:         "NetWork Test",
            Shared:       false,
            AdminStateUp: true,
            TenantId:     "57e98940a77f4bb988a21a7d0603a626",
        },
    }).Extract()

    if err != nil {
        panic(err)
    }

Sample Code, Updates a network.


    result, err := networks.Update(client, "1c6af92e-bd06-4350-8f51-5ec32167814f", networks.UpdateOpts{
        Network: networks.UpdateNetwork{
            Name:   "Test Net Works Updated",
            Shared: true,
        },
    }).Extract()

    if err != nil {
        panic(err)
    }

Sample Code, Queries details about the specified network.


    result, err := networks.Get(client, "1c6af92e-bd06-4350-8f51-5ec32167814f").Extract()

    if err != nil {
        panic(err)
    }

Sample Code, Deletes all networks to which the specified tenant has access.


    result := networks.Delete(client, "1c6af92e-bd06-4350-8f51-5ec32167814f")

Sample Code, Queries all networks accessible to the tenant submitting the request. A maximum of 2000 records can be returned for each query operation.


    allPages, err := networks.List(client, networks.ListOpts{
        Limit: 2,
    }).AllPages()

    result, err := networks.ExtractList(allPages.(networks.ListPage))

    if err != nil {
        panic(err)
    }
*/
package networks
