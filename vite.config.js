import { defineConfig } from "vite";
import { globSync } from "glob";
import { resolve } from "path";
import { viteStaticCopy } from 'vite-plugin-static-copy';
import  viteCompression  from "vite-plugin-compression";
import tailwindcss from '@tailwindcss/vite';

const root   = resolve(__dirname, '.');
// static data
const publicDir  = resolve(__dirname, 'web/public');
// build targets
const input = globSync(
    [
        "web/src/index.{js,css}",
        "web/src/**/*.{js,css}",
    ]).map(
        (path) => resolve(process.cwd(), path)
    );

export default defineConfig({
    plugins: [
        tailwindcss(),
        viteStaticCopy({
            targets: [
                {
                    src: "node_modules/material-symbols/outlined.css",
                    dest: 'fonts',
                    rename: 'material-symbols-outlined.css'
                },
                {
                    src: 'node_modules/material-symbols/material-symbols-outlined.woff2',
                    dest: 'fonts'
                }
            ]
        }),
       viteCompression({
            algorithm: "brotliCompress",
            ext: ".br",
            threshold: 10240,
            deleteOriginFile: false,
        }),
    ],
    root,
    publicDir,
    // entry: {core: resolve(root, "/src/main.js")},
    build: {
        // outDir: resolve(process.cwd(), "dist"),
        outDir: "./dist",
        assetsDir: "assets",
        emptyOutDir: true,
        sourcemap: true,
        rollupOptions: {
            input,
            output: {
                entryFileNames: '[name].js',
                chunkFileNames: 'assets/[name]-chunk.js',
                assetFileNames: 'assets/[name][extname]',
            },
        },
    },
})
