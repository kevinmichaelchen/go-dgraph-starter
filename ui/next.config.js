const isDev = process.env.NODE_ENV === "development";

let i18n = {
  // These are all the locales you want to support in
  // your application
  locales: ["en", "es", "fr"],
  // This is the default locale you want to be used when visiting
  // a non-locale prefixed path e.g. `/hello`
  defaultLocale: "en",
};

if (!isDev) {
  i18n = {
    ...i18n,
    // This is a list of locale domains and the default locale they
    // should handle (these are only required when setting up domain routing)
    // Note: subdomains must be included in the domain value to be matched e.g. "fr.example.com".
    domains: [
      {
        domain: "example.com",
        defaultLocale: "en",
      },
      {
        domain: "es.example.com",
        defaultLocale: "es",
      },
      {
        domain: "fr.example.com",
        defaultLocale: "fr",
      },
    ],
  };
}

module.exports = {
  images: {
    domains: ["example.com", "upload.wikimedia.org"],
  },
  i18n,
};
