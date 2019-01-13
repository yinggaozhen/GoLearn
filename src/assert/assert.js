const assert = require('assert');

/**
 * assert(value[, message])
 *
 * @link  : http://nodejs.cn/api/assert.html#assert_assert_value_message
 * @since : v0.5.9
 */
class Assert {
    static run() {
        assert(1, 'error');
        assert(0, 'invalid true value');
    }
}

Assert.run();