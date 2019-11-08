var { AES, enc } = require('crypto-js');

/**
 *
 * Awise socket client
 * @param {string} uri
 * @version 1.0.0
 * @example AwiseSocket('wss://messenger.awise.co')
 * @author Nyura95
 *
 */
function AwiseSocket(uri, logger = true) {
  if (typeof uri !== 'string' && (uri.indexOf('ws://') === -1 || uri.indexOf('wss://') === -1)) {
    throw 'This uri is not correct, you must pass a ws or wss uniform resource identifier';
  }
  Object.defineProperty(this, '_uri', {
    value: uri,
    writable: false,
  });
  this._tokenConversation = null;
  this._cryptKey = null;
  this.webSocket = null;
  this.logger = logger;
  this.onerror = function() {};
  this.onclose = function() {};

  // private action
  this.message = function() {};
  this.update = function() {};
  this.delete = function() {};
  this.private = function() {};
  this.connection = function() {};
  this.disconnection = function() {};
  this.error = function() {};
}

/**
 *
 * Initialize a new connexion with a conversation
 * @param {string} token
 * @param {function} callback
 * @version 1.0.0
 * @returns {void}
 * @author Nyura95
 */
AwiseSocket.prototype.initConversation = function(token, callback) {
  if (this.webSocket && this.webSocket.readyState === WebSocket.OPEN) {
    this.close();
  }
  this._tokenConversation = token;
  this._log('init conversation');
  this.webSocket = new WebSocket(this._uri + '/' + this._tokenConversation);
  this.webSocket.onopen = function() {
    callback();
  };
  this.webSocket.onerror = function(event) {
    this.onerror ? this.onerror(event) : null;
  }.bind(this);
  this.webSocket.onclose = function(event) {
    this.onclose ? this.onclose(event) : null;
  }.bind(this);
  this.webSocket.onmessage = function(event) {
    var actions = event.data.split('\n');
    for (let i = 0; i < actions.length; i++) {
      var action = '';
      try {
        action = JSON.parse(actions[i]);
      } catch (err) {
        console.log(actions);
        this._log('error parsing actions');
        break;
      }
      this._log('new message receive (' + action.action + ')');
      console.log(action);
      if (action.action === 'message') {
        if (this._cryptKey) {
          action.message.message = this._decrypt(action.message.message);
        }
        this.message(action.message);
      }
      if (action.action === 'update') {
        this.update(action.message);
      }
      if (action.action === 'delete') {
        this.delete(action.message);
      }
      if (action.action === 'connection') {
        this.connection(action.user);
      }
      if (action.action === 'disconnection') {
        this.disconnection(action.user);
      }
      if (action.action === 'private') {
        this.privateMode(action.token);
      }
      if (action.action === 'error') {
        this.error(action.locKey, action.message);
      }
    }
  }.bind(this);
};

/**
 *
 * Send a message to the conversation
 * @param {string} token
 * @version 1.0.0
 * @returns {void}
 * @author Nyura95
 */
AwiseSocket.prototype.send = function(message) {
  if (!this._tokenConversation && typeof this._tokenConversation !== 'number') {
    console.warn('You must target a conversation (use initConversation)');
    return;
  }
  this._log('send message : (' + message + ')');
  this.webSocket ? this.webSocket.send(this._cryptKey ? this._encrypt(message) : message) : null;
};

/**
 *
 * Close the server connexion
 * @version 1.0.0
 * @returns {void}
 * @author Nyura95
 */
AwiseSocket.prototype.close = function() {
  this._log('close');
  this.webSocket.close(1000, 'close user');
  this.webSocket = null;
};

/**
 *
 * Activate the private mode
 * @param {string} token
 * @version 1.0.0
 * @returns {void}
 * @author Nyura95
 */
AwiseSocket.prototype.privateMode = function(cryptKey) {
  this._log('pivate mode');
  if (confirm('Voulez-vous activer le private mode sur cette conversation ?')) {
    this._cryptKey = cryptKey;
    this.PrivateMode = true;
    this.private(cryptKey);
  }
};

/**
 *
 * Activate the public mode
 * @version 1.0.0
 * @returns {void}
 * @author Nyura95
 */
AwiseSocket.prototype.publicMode = function() {
  this._log('public mode');
  this._cryptKey = null;
  this.PrivateMode = false;
};

/**
 *
 * Return if the private mode is active
 * @version 1.0.0
 * @author Nyura95
 */
AwiseSocket.prototype.PrivateMode = false;

/**
 *
 * Encrypt a message
 * @param {string} message
 * @version 1.0.0
 * @returns {void}
 * @author Nyura95
 * @private
 */
AwiseSocket.prototype._encrypt = function(message) {
  this._log('encrypt message');
  if (!this._tokenConversation || !this._cryptKey) {
    console.warn('You must target a conversation and activate a private mode');
    return null;
  }
  return AES.encrypt(message, this._cryptKey).toString();
};

/**
 *
 * Decrypt a message
 * @param {string} message
 * @version 1.0.0
 * @returns {void}
 * @author Nyura95
 * @private
 */
AwiseSocket.prototype._decrypt = function(hash) {
  this._log('decrypt message');
  if (!this._tokenConversation || !this._cryptKey) {
    console.warn('You must target a conversation and activate a private mode');
    return null;
  }
  return AES.decrypt(hash, this._cryptKey).toString(enc.Utf8);
};

/**
 *
 * Log socket
 * @param {Array<string>} messages
 * @version 1.0.0
 * @returns {void}
 * @author Nyura95
 * @private
 */
AwiseSocket.prototype._log = function(...messages) {
  if (console) {
    console.log(`[${this._tokenConversation || 'noToken'}]:`, ...messages);
  }
};

module.exports = AwiseSocket;

/// <reference path=”./index.d.ts” />
