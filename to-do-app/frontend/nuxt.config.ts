// nuxt.config.ts

export default defineNuxtConfig({
  compatibilityDate: '2024-04-03',
  devtools: { enabled: true },
  app: {
    // head: {
    //   link: [{ rel: 'stylesheet', href: '/assets/css/styles.css' }],
    //   script: [
        // { src: '/assets/libs/bootstrap/dist/js/bootstrap.bundle.min.js', type: 'text/javascript' },
        // { src: '/assets/libs/simplebar/dist/simplebar.min.js', type: 'text/javascript' },
        // { src: '/assets/js/theme/app.init.js', type: 'text/javascript' },
        // { src: '/assets/js/theme/theme.js', type: 'text/javascript' },
        // { src: '/assets/js/theme/app.min.js', type: 'text/javascript' }
    //   ]
    // },
    head: {
      link: [
        { rel: 'stylesheet', href: '/assets/css/09f6b43cb9566376.css', 'data-precedence': 'next' },
        { rel: 'stylesheet', href: '/assets/css/e30dffe59c84f146.css', 'data-precedence': 'next' },
        { rel: 'stylesheet', href: '/assets/css/97924461c38f61c9.css', 'data-precedence': 'next' },
        { rel: 'stylesheet', href: '/assets/css/c3c1fbb4abd379f9.css', 'data-precedence': 'next' },
        { rel: 'stylesheet', href: '/assets/css/styles.css' }
      ],
      script: [
        { src: '/assets/libs/bootstrap/dist/js/bootstrap.bundle.min.js', type: 'text/javascript' },
        { src: '/assets/libs/simplebar/dist/simplebar.min.js', type: 'text/javascript' },
        { src: '/assets/js/theme/app.init.js', type: 'text/javascript' },
        { src: '/assets/js/theme/theme.js', type: 'text/javascript' },
        { src: '/assets/js/theme/app.min.js', type: 'text/javascript' },
        { src: '/assets/js/fd9d1056-b9b3cb082750b374.js', async: true },
        { src: '/assets/js/7023-c86d20109cd6fc48.js', async: true },
        { src: '/assets/js/main-app-b43adc49a3c065b6.js', async: true },
        { src: '/assets/js/795d4814-2e40a0ec8a548560.js', async: true },
        { src: '/assets/js/53c13509-fd04a802f83abbfc.js', async: true },
        { src: '/assets/js/8e1d74a4-fb253ec53dd3ba72.js', async: true },
        { src: '/assets/js/66ec4792-48a4fbd036915e69.js', async: true },
        { src: '/assets/js/d7101eed-4d705c283a471601.js', async: true },
        { src: '/assets/js/e8686b1f-a673c053cc68b8b1.js', async: true },
        { src: '/assets/js/7301-bb09c558bc2c714f.js', async: true },
        { src: '/assets/js/8173-02c44c13c7f64877.js', async: true },
        { src: '/assets/js/5506-83fe3495650aa4a5.js', async: true },
        { src: '/assets/js/9355-720912148d3dd789.js', async: true },
        { src: '/assets/js/9109-26e9436acfa055bb.js', async: true },
        { src: '/assets/js/3208-d42a0d07d5a5f71d.js', async: true },
        { src: '/assets/js/3151-626447b236980efa.js', async: true },
        { src: '/assets/js/8446-d8061d8b8b435bcc.js', async: true },
        { src: '/assets/js/3139-371471cf0f4d7e90.js', async: true },
        { src: '/assets/js/5621-59194c5afa0b4ab2.js', async: true },
        { src: '/assets/js/page-7c38cbcf65addbd2.js', async: true },
        { src: '/assets/js/231-95ebbb80c4cff189.js', async: true },
        { src: '/assets/js/6449-15305117ff3ac449.js', async: true },
        { src: '/assets/js/6162-ccd3f985a81a9b7d.js', async: true },
        { src: '/assets/js/9224-1628b945caba1d43.js', async: true },
        { src: '/assets/js/layout-bfd575b834f3ce68.js', async: true },
        { src: '/assets/js/69c8c2c4-fd5a09277478a795.js', async: true },
        { src: '/assets/js/6378-2b7fb941c473786e.js', async: true },
        { src: '/assets/js/8995-4a38ae5bea613153.js', async: true },
        { src: '/assets/js/layout-35860d29da3ab64d.js', async: true },
        { src: '/assets/js/not-found-b1741c041364087b.js', async: true }
      ]
    }
  }
});
