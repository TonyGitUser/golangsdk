/*


Sample Code, This interface is used to query the currently supported API version list.


    result, err := versions.List(client).Extract()

    if err != nil {
        panic(err)
    }

Sample Code, This interface is used to query the specified API version.


    result, err := versions.Get(client, "v1").Extract()

    if err != nil {
        panic(err)
    }
*/
package versions
