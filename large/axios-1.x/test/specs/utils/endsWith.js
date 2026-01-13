import utils from '../../../lib/utils';

const {kindOf} = utils;

describe('utils::kindOf', function () {
  it('should return object tag', function () {
    expect(kindOf({})).toEqual('object');
    // cached result
    expect(kindOf({})).toEqual('object');
    expect(kindOf([])).toEqual('array');
  });
});
// ID-1768294468-27a1a742
