fetch(
  'https://api.worldoftanks.asia/wot/encyclopedia/vehicles/?application_id=d5c27b088716f6a2ca4d043e6fe2ba91&tier=1&language=vi&page_no=1',
  {
    headers: {
      accept:
        'text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7',
      'accept-language': 'en-US,en;q=0.9,vi;q=0.8,ko;q=0.7',
      'cache-control': 'no-cache',
      pragma: 'no-cache',
      'sec-ch-ua':
        '"Chromium";v="140", "Not=A?Brand";v="24", "Google Chrome";v="140"',
      'sec-ch-ua-mobile': '?0',
      'sec-ch-ua-platform': '"Windows"',
      'sec-fetch-dest': 'document',
      'sec-fetch-mode': 'navigate',
      'sec-fetch-site': 'none',
      'sec-fetch-user': '?1',
      'upgrade-insecure-requests': '1',
    },
    body: null,
    method: 'GET',
    mode: 'cors',
    credentials: 'include',
  }
)
  .then((a) => a.json())
  .then((b) => console.log('b', b));
