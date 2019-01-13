const assert = require('assert');

/**
 * assert.deepEqual(actual, expected[, message])
 *
 * @link  : http://nodejs.cn/api/assert.html#assert_assert_deepequal_actual_expected_message
 * @since : v0.5.9
 */
class AssertDeepEqual {
    static run() {
        const obj1 = {
            a: {
                b: 1
            }
        };
        const obj2 = {
            a: {
                b: 2
            }
        };
        const obj3 = {
            a: {
                b: 1
            }
        };
        const obj4 = {
            a: {
                b: '1'
            }
        };
        const obj5 = Object.create(obj1);

        // OK
        assert.deepEqual(obj1, obj1);

        // Values of b are different:
        // AssertionError: { a: { b: 1 } } deepEqual { a: { b: 2 } }
        // assert.deepEqual(obj1, obj2);

        // OK
        assert.deepEqual(obj1, obj3);

        // Prototypes are ignored:
        // AssertionError: { a: { b: 1 } } deepEqual {}
        assert.deepEqual(obj1, obj4);

        // Prototypes are ignored:
        // AssertionError: { a: { b: 1 } } deepEqual {}
        assert.deepEqual(obj1, obj5);
    }
}

AssertDeepEqual.run();