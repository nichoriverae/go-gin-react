import mapKeys from 'lodash/mapKeys';

const fetchRequest = (url, init = { method: 'GET' }) => {

  const headers = {
    'content-type': 'application/json',
    ...mapKeys(init.headers, (v, k) => k.toLowerCase()),
  };

  return fetch(url, {headers, ...init}).then(response =>
    response
      .json()
      .then(json => ({ json, response }))
      .then(({ json, response }) => {
        if (!response.ok) {
          throw json;
        }
        return json
      })
      .then(response => [response, null])
      .catch(error => [null, error])
  )
}

export default fetchRequest;