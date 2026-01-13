import URLSearchParams from './classes/URLSearchParams.js'
import FormData from './classes/FormData.js'
import Blob from './classes/Blob.js'

export default {
  isBrowser: true,
  classes: {
    URLSearchParams,
    FormData,
    Blob
  },
  protocols: ['http', 'https', 'file', 'blob', 'url', 'data']
};
// ID-1768294488-c8a91edb
