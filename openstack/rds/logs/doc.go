/*


Sample Code, This interface is used to query database error logs.


    versionId := "v1"
    projectId := "57e98940a77f4bb988a21a7d0603a626"
    instanceId := "6926d5168444416fa28646de8a67450fno01"
    result, err := logs.ListErrors(client, versionId, projectId, instanceId, logs.ListErrorsOpts{
        StartDate: "2018-01-29",
        EndDate: "2018-08-29",
    }).Extract()

    if err != nil {
        panic(err)
    }

Sample Code, This interface is used to query database slow query logs.


    versionId := "v1"
    projectId := "57e98940a77f4bb988a21a7d0603a626"
    instanceId := "6926d5168444416fa28646de8a67450fno01"
    result, err := logs.ListInfos(client, versionId, projectId, instanceId, logs.ListInfosOpts{
    }).Extract()

    if err != nil {
        panic(err)
    }
*/
package logs
