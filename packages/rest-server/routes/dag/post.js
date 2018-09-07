'use strict'

const {
  Joi,
} = require('../../utils/validation')

module.exports = {
  method: 'POST',
  path: '/dag/{cid}',
  options: {
    handler: (request, reply) => {},
    description: 'Create a DAG node',
    notes: 'Resolves a DAG node with the passed CID',
    tags: ['api'],
    validate: {
      params: {
        cid: Joi
          .cid()
          .description('The CID that corresponds to the DAG node we wish to create')
      }
    },
    plugins: {
      'hapi-swagger': {
        id: 'create'
      }
    }
  }
}
