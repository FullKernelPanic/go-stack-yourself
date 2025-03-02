import * as esbuild from 'esbuild';
import {sassPlugin} from 'esbuild-sass-plugin';

const buildOptions = {
    entryPoints: ['js/main.js', "scss/main.scss"],
    plugins: [sassPlugin()],
    bundle: true,
    outdir: '../public/',
    minify: true,
    sourcemap: true,
    target: ['es2020'],
}

// Check for "watch" flag, e.g., `node build.js --watch`
const args = process.argv.slice(2);
const isWatchMode = args.includes('--watch');


(async () => {
    if (isWatchMode) {
        // Use esbuild's `context()` for watch mode
        const ctx = await esbuild.context(buildOptions);

        // Start watch mode
        await ctx.watch();
        console.log('Watching for changes...');
    } else {
        // Normal build mode (no watch)
        esbuild.build(buildOptions).then(() => {
            console.log('Build succeeded!');
        }).catch((e) => {
            console.error('Build failed:', e.message);
            process.exit(1);
        });
    }
})();