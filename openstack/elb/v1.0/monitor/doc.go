/*
This interface is used to query all layer-4 and layer-7 traffic monitoring metrics.Only common tenants can query these monitoring metrics.

Sample Code, This interface is used to query all layer-4 and layer-7 traffic monitoring metrics.Only common tenants can query these monitoring metrics.


    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result, err := monitor.List(client, tenantID).Extract()

    if err != nil {
        panic(err)
    }
*/
package monitor
