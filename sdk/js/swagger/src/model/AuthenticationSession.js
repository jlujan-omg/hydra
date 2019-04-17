/**
 * ORY Hydra
 * Welcome to the ORY Hydra HTTP API documentation. You will find documentation for all HTTP APIs here.
 *
 * OpenAPI spec version: latest
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.2.3
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'));
  } else {
    // Browser globals (root is window)
    if (!root.OryHydra) {
      root.OryHydra = {};
    }
    root.OryHydra.AuthenticationSession = factory(root.OryHydra.ApiClient);
  }
}(this, function(ApiClient) {
  'use strict';




  /**
   * The AuthenticationSession model module.
   * @module model/AuthenticationSession
   * @version latest
   */

  /**
   * Constructs a new <code>AuthenticationSession</code>.
   * AuthenticationSession AuthenticationSession AuthenticationSession AuthenticationSession AuthenticationSession AuthenticationSession AuthenticationSession AuthenticationSession AuthenticationSession AuthenticationSession AuthenticationSession AuthenticationSession AuthenticationSession AuthenticationSession AuthenticationSession authentication session
   * @alias module:model/AuthenticationSession
   * @class
   */
  var exports = function() {
    var _this = this;




  };

  /**
   * Constructs a <code>AuthenticationSession</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/AuthenticationSession} obj Optional instance to populate.
   * @return {module:model/AuthenticationSession} The populated <code>AuthenticationSession</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('AuthenticatedAt')) {
        obj['AuthenticatedAt'] = ApiClient.convertToType(data['AuthenticatedAt'], 'Date');
      }
      if (data.hasOwnProperty('ID')) {
        obj['ID'] = ApiClient.convertToType(data['ID'], 'String');
      }
      if (data.hasOwnProperty('Subject')) {
        obj['Subject'] = ApiClient.convertToType(data['Subject'], 'String');
      }
    }
    return obj;
  }

  /**
   * authenticated at Format: date-time Format: date-time Format: date-time Format: date-time Format: date-time Format: date-time Format: date-time Format: date-time Format: date-time Format: date-time Format: date-time Format: date-time Format: date-time Format: date-time Format: date-time
   * @member {Date} AuthenticatedAt
   */
  exports.prototype['AuthenticatedAt'] = undefined;
  /**
   * ID
   * @member {String} ID
   */
  exports.prototype['ID'] = undefined;
  /**
   * subject
   * @member {String} Subject
   */
  exports.prototype['Subject'] = undefined;



  return exports;
}));


