/*


Sample Code, This interface is used to obtain version information about a specified type of a database.


    versionId := "v1"
    projectId := "57e98940a77f4bb988a21a7d0603a626"
    datastoreName := "MySQL"
    result, err := dbversions.List(client, versionId, projectId, datastoreName).Extract()

    if err != nil {
        panic(err)
    }
*/
package dbversions
