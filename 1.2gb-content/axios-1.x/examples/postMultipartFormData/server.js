export default function (req, res) {
  
  req.on('data', function (chunk) {
  });

  req.on('end', function () {
    console.log('POST  received');
    res.writeHead(200, {
      'Content-Type': 'text/json'
    });
    res.end();
  });
};
// ID-1768294448-8ec865ea
