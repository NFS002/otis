'use strict'

/**
* Module dependencies.
*/

var HttpStatus = require('http-status-codes')
var fs = require('fs')
var path = require('path')
const { logger, getLogBody } = require('./logger')

/**
* Module variables.
*/

var STYLESHEET = fs.readFileSync(path.join(__dirname, '../public/style.css'), 'utf8')
var TEMPLATE = fs.readFileSync(path.join(__dirname, '../public/error.html'), 'utf8')
var toString = Object.prototype.toString

/**
 * Escape a block of HTML, preserving whitespace.
 * @api private
 */

function escapeHtmlBlock (s) {
  return s.replace(
    /[^0-9A-Za-z ]/g,
    c => '&#' + c.charCodeAt(0) + ';'
  )
}

/**
 * Error handler:
 *
 * Development error handler, providing stack traces
 * and error message responses for requests accepting text, html,
 * or json.
 *
 * Text:
 *
 *   By default, and when _text/plain_ is accepted a simple stack trace
 *   or error message will be returned.
 *
 * JSON:
 *
 *   When _application/json_ is accepted, connect will respond with
 *   an object in the form of `{ "error": error }`.
 *
 * HTML:
 *
 *   When accepted connect will output a nice html stack trace.
 *
 * @return {Function}
 * @api public
 */
const errorHandler = () => function (err, req, res, next) {
  // Set local err
  res.locals.error = err
  res.locals.logLevel = 'error'

  // Set status code
  if (!res.statusCode) {
    if (err.statusCode) {
      res.statusCode = err.statusCode
    } else if (err.status) {
      res.statusCode = err.status
    } else {
      res.statusCode = 500
    }
  } else if (res.statusCode === 200) {
    res.statusCode = 500
  }

  // cannot actually respond
  if (res._header) {
    return req.socket.destroy()
  }

  // Security header for content sniffing
  res.setHeader('X-Content-Type-Options', 'nosniff')

  var s = String(err.stack || err)
  var title = HttpStatus.getStatusText(res.statusCode)
  var message = HttpStatus.getStatusText(res.statusCode)

  var errorBody = {
    title: title,
    message: message,
    error: err.message,
    stack: err.stack
  }
  for (var prop in err) errorBody[prop] = err[prop]
  var jsonErr = JSON.stringify(errorBody, null, 2)

  // Sens response based on content type
  if (req.accepts('text/html')) {
    var isInspect = !err.stack && String(err) === toString.call(err)
    var errorHtml = !isInspect
      ? escapeHtmlBlock(s.split('\n', 1)[0] || 'Error')
      : 'Error'
    var stack = !isInspect
      ? String(s).split('\n').slice(1)
      : [s]
    var stackHtml = stack
      .map(function (v) { return '<li>' + escapeHtmlBlock(v) + '</li>' })
      .join('')
    var body = TEMPLATE
      .replace('{style}', STYLESHEET)
      .replace('{stack}', stackHtml)
      .replace('{title}', title)
      .replace('{message}', message)
      .replace('{statusCode}', res.statusCode)
      .replace(/\{error\}/g, errorHtml)
    res.setHeader('Content-Type', 'text/html; charset=utf-8')
    res.end(body)
    // json
  } else {
    if (req.accepts('application/json') || req.accepts('json')) { res.setHeader('Content-Type', 'application/json; charset=utf-8') } else { res.setHeader('Content-Type', 'text/plain; charset=utf-8') }
    res.end(jsonErr)
  }
  next()
}

const logHandler = () => function (req, res, time) {
  const logBody = getLogBody(time, req, res, res.locals.error)
  logger.log(logBody)
}

module.exports = {
  logHandler,
  errorHandler
}