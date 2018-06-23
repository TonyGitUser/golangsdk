/*
This section describes how to manage HTTPS certificates. You can upload a certificate and bind it to an HTTPS listener to provide the HTTPS or TCP service.

Sample Code, This interface is used to create a certificate for HTTPS.


    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result, err := certificate.Create(client, tenantID, certificate.CreateOpts{
        Name:        "Cert1",
        Description: "Testing Cert1",
        Domain:      "",
        Certificate: `-----BEGIN CERTIFICATE-----
    MIIDIjCCAougAwIBAgIJALV96mEtVF4EMA0GCSqGSIb3DQEBBQUAMGoxCzAJBgNVBAYTAnh4MQswCQYDVQQIEwJ4eDELMAkGA1UEBxMCeHgxCzAJBgNVBAoTAnh4MQswCQYDVQQLEwJ4eDELMAkGA1UEAxMCeHgxGjAYBgkqhkiG9w0BCQEWC3h4eEAxNjMuY29tMB4XDTE3MTExMzAyMjYxM1oXDTIwMTExMjAyMjYxM1owajELMAkGA1UEBhMCeHgxCzAJBgNVBAgTAnh4MQswCQYDVQQHEwJ4eDELMAkGA1UEChMCeHgxCzAJBgNVBAsTAnh4MQswCQYDVQQDEwJ4eDEaMBgGCSqGSIb3DQEJARYLeHh4QDE2My5jb20wgZ8wDQYJKoZIhvcNAQEBBQADgY0AMIGJAoGBAMU832iM+d3FILgTWmpZBUoYcIWVcAAYE7FsZ9LNerOyjJpyi256oypdBvGs9JAUBN5WaFk81UQx29wAyNixX+bKa0DBWpUDqr84V1f9vdQc75v9WoujcnlKszzpV6qePPC7igJJpu4QOI362BrWzJCYQbg4Uzo1KYBhLFxl0TovAgMBAAGjgc8wgcwwHQYDVR0OBBYEFMbTvDyvE2KsRy9zPq/JWOjovG+WMIGcBgNVHSMEgZQwgZGAFMbTvDyvE2KsRy9zPq/JWOjovG+WoW6kbDBqMQswCQYDVQQGEwJ4eDELMAkGA1UECBMCeHgxCzAJBgNVBAcTAnh4MQswCQYDVQQKEwJ4eDELMAkGA1UECxMCeHgxCzAJBgNVBAMTAnh4MRowGAYJKoZIhvcNAQkBFgt4eHhAMTYzLmNvbYIJALV96mEtVF4EMAwGA1UdEwQFMAMBAf8wDQYJKoZIhvcNAQEFBQADgYEAASkC/1iwiALa2RU3YCxqZFEEsZZvQxikrDkDbFeoa6Tk49Fnb1f7FCW6PTtY3HPWl5ygsMsSy0Fi3xp3jmuIwzJhcQ3tcK5gC99HWp6Kw37RL8WoB8GWFU0Q4tHLOjBIxkZROPRhH+zMIrqUexv6fsb3NWKhnlfh1Mj5wQE4Ldo=
    -----END CERTIFICATE-----`,
        PrivateKey: `-----BEGIN RSA PRIVATE KEY-----
    MIICXQIBAAKBgQDFPN9ojPndxSC4E1pqWQVKGHCFlXAAGBOxbGfSzXqzsoyacotueqMqXQbxrPSQFATeVmhZPNVEMdvcAMjYsV/mymtAwVqVA6q/OFdX/b3UHO+b/VqLo3J5SrM86Veqnjzwu4oCSabuEDiN+tga1syQmEG4OFM6NSmAYSxcZdE6LwIDAQABAoGBAJvLzJCyIsCJcKHWL6onbSUtDtyFwPViD1QrVAtQYabF14g8CGUZG/9fgheuTXPtTDcvu7cZdUArvgYW3I9F9IBb2lmF3a44xfiAKdDhzr4DK/vQhvHPuuTeZA41r2zp8Cu+Bp40pSxmoAOK3B0/peZAka01Ju7c7ZChDWrxleHZAkEA/6dcaWHotfGSeW5YLbSms3f0m0GH38nRl7oxyCW6yMIDkFHURVMBKW1OhrcuGo8u0nTMi5IH9gRg5bH8XcujlQJBAMWBQgzCHyoSeryD3TFieXIFzgDBw6Ve5hyMjUtjvgdVKoxRPvpOkclc39QHP6Dm2wrXXHEej+9RILxBZCVQNbMCQQC42i+Ut0nHvPuXN/UkXzomDHdeh1ySsOAO4H+8Y6OSI87l3HUrByCQ7stX1z3L0HofjHqV9Koy9emGTFLZEzSdAkB7Ei6cUKKmztkYe3rr+RcATEmwAw3tEJOHmrW5ErApVZKr2TzLMQZ7WZpIPzQRCYnY2ZZLDuZWFFG3vW+wKKktAkAaQ5GNzbwkRLpXF1FZFuNF7erxypzstbUmU/31b7tSi5LmxTGKL/xRYtZEHjya4Ikkkgt40q1MrUsgIYbFYMf2
    -----END RSA PRIVATE KEY-----`,
    }).Extract()

    if err != nil {
        panic(err)
    }

Sample Code, This interface is used to delete a certificate.


    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    certificateId := "3d42b568bb2147f7a241c6a07dc3715d"
    certificate.Delete(client, tenantID, certificateId)

Sample Code, This interface is used to query the certificate list.


    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result, err := certificate.List(client, tenantID).Extract()

    if err != nil {
        panic(err)
    }

Sample Code, This interface is used to change the name and description of a certificate.


    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    certificateId := "3d42b568bb2147f7a241c6a07dc3715d"
    result, err := certificate.Update(client, tenantID, certificateId, certificate.UpdateOpts{
        Name:        "UpdatedCert1",
        Description: "TEST ABC",
    }).Extract()

    if err != nil {
        panic(err)
    }
*/
package certificate
