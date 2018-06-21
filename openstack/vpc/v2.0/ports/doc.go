/*
This interface is used to manage and perform operations on ports, including querying ports, creating a port, querying a specified port, deleting a port, and updating a port.

Sample Code, This interface is used to create a port.


    result, err := ports.Create(client, ports.CreateOpts{
        Port: ports.CreatePort{
            Name:      "EricTestPort",
            NetworkId: "5ae24488-454f-499c-86c4-c0355704005d",
        },
    }).Extract()

    if err != nil {
        panic(err)
    }

Sample Code, This interface is used to update a port.


    result, err := ports.Update(client, "d6d0258f-7bf2-4504-845b-ab6b8513a986", ports.UpdateOpts{
        Port: ports.UpdatePort{
            Name: "ModifiedPort",
        },
    }).Extract()

    if err != nil {
        panic(err)
    }

Sample Code, This interface is used to query a single port.


    result, err := ports.Get(client, "d6d0258f-7bf2-4504-845b-ab6b8513a986").Extract()

    if err != nil {
        panic(err)
    }

Sample Code, This interface is used to query ports and to display the ports in a list.


    allPages, err := ports.List(client, ports.ListOpts {
        Limit: 2,
        Name: "EricTestPort",
    }).AllPages()

    result, err := ports.ExtractList(allPages.(ports.ListPage))

    if err != nil {
        panic(err)
    }

Sample Code, This interface is used to delete a port.

    result := ports.Delete(client, "d6d0258f-7bf2-4504-845b-ab6b8513a986")

    if err != nil {
        panic(err)
    }
*/
package ports
