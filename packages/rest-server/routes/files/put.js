'use strict'

const Joi = require('joi')

module.exports = {
  method: 'PUT',
  path: '/files/{path}',
  config: {
    handler: (request, reply) => {},
    description: 'Update an MFS path',
    notes: 'Replace a file or directory at the passed MFS path',
    tags: ['api'],
    validate: {
      params: {
        path: Joi
          .string()
          .required()
          .description('The MFS path you wish to replace')
      }
    }
  }
}
