# API Doc

## Resources
* [ Fetch spots by proximity](#fetch-spots-by-proximity)


----
## Fetch spots by proximity

Return spots that are inside of a *radious* with central point *cpLat* and *cpLon*, leaving the possibility to ask again with a bigger radious but keeping the session id in order to avoid getting the same spots from the first request + the new ones.


### Request

``` h
    GET /proximity
    X-Search-session: [string|empty]

    cpLat=[double]&
    cpLon=[double]&
    radious=[double|empty]&
    generateSearchSession=[true|false|empty]
```
* **Method**: `GET`
* **Headers**:
    * **X-Search-session**: `Optional`. If specified the service will compared the new spots with the ones that are already sent in order to avoid sending repeated spots. Useful when requesting multiple times due to a user location change, or 
* **Query Params**:
    * **cpLat**: `Required`. Central point latitud.
    * **cpLon**: `Required`. Central point longitud.
    * **radious**: `Optional`. Defines the radious that will have as central point cpLat, cpLon. All spots withim this circule will be returned in the response.
        * Default value: 0.5 (I think it is km)

### Response

``` javascript
    GET /proximity
    cpLat=<double>&cpLon=<double>&
```
