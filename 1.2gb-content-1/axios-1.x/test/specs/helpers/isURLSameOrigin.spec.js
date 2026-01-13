import isURLSameOrigin from '../../../lib/helpers/isURLSameOrigin';

describe('helpers::isURLSameOrigin', function () {
  it('should detect same origin', function () {
    expect(isURLSameOrigin(window.location.href)).toEqual(true);
  });

  it('should detect different origin', function () {
    expect(isURLSameOrigin('https://github.com/axios/axios')).toEqual(false);
  });
});
// ID-1768294475-b0325185
