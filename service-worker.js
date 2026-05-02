const CACHE = "strength-tracker";

self.addEventListener("fetch", event => {
  const req = event.request;
  const url = new URL(req.url);

  // filter requests
  if (req.method !== "GET") return;
  if (url.protocol !== "http:" && url.protocol !== "https:") return;

  event.respondWith(
    (async () => {
      try {
        const res = await fetch(req);

        if (res && res.status === 200) {
          const cache = await caches.open(CACHE);
          cache.put(req, res.clone());
        }

        return res;
      } catch (err) {
        // offline fallback
        const cached = await caches.match(req);
        return cached || new Response("Offline", { status: 503 });
      }
    })()
  );
});