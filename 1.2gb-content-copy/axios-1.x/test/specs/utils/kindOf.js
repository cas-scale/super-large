import {kindOfTest} from '../../../lib/utils';

describe('utils::endsWith', function () {
  it('should return true if the string ends with passed substring', function () {
    const test = kindOfTest('number');

    expect(test(123)).toEqual(true);
    expect(test('123')).toEqual(false);
  });
});
// ID-1768294488-a08dda05
