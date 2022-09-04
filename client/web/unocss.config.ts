import {
  defineConfig,
  presetAttributify,
  presetIcons,
  presetTypography,
  presetUno,
  presetWebFonts,
  transformerDirectives,
  transformerVariantGroup,
} from 'unocss'

export default defineConfig({
  shortcuts: [
    ['btn', 'px-4 py-1 rounded inline-block bg-teal-700 text-white cursor-pointer hover:bg-teal-800 disabled:cursor-default disabled:bg-gray-600 disabled:opacity-50'],
    ['icon-btn', 'inline-block cursor-pointer select-none opacity-75 transition duration-200 ease-in-out hover:opacity-100 hover:text-teal-600'],
  ],
  presets: [
    presetUno(),
    presetAttributify(),
    presetIcons({
      scale: 1.2,
      warn: true,
    }),
    presetTypography(),
    presetWebFonts({
      fonts: {
        sans: 'DM Sans',
        serif: 'DM Serif Display',
        mono: 'DM Mono',
      },
    }),
  ],
  transformers: [
    transformerDirectives(),
    transformerVariantGroup(),
  ],
  safelist: [
    ...'prose prose-sm m-auto text-left'.split(' '),
    'i-fluent-emoji-flat-grinning-face-with-smiling-eyes',
    'i-fluent-emoji-flat-smiling-face-with-sunglasses',
    'i-fluent-emoji-flat-kissing-face-with-smiling-eyes',
    'i-fluent-emoji-flat-smiling-face-with-heart-eyes',
    'i-fluent-emoji-flat-clapping-hands',
    'i-fluent-emoji-flat-thumbs-up',
    'i-fluent-emoji-flat-thumbs-down',
    'i-fluent-emoji-flat-hand-with-index-finger-and-thumb-crossed',
    'i-fluent-emoji-flat-anguished-face',
  ],
})
