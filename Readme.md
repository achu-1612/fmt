# ~~FMT~~ Diff

Oh, this project? Yeah, it was totally born from a bout of severe boredom and probably too much caffeine. I mean, who doesn't randomly decide to build a text comparison tool just for kicks?

So here’s what it does: you throw in two chunks of text, and it highlights every little nitpicky difference between them. Like, maybe you added an extra “and” or a sneaky typo—it’s catching it! It's powered by Go and WebAssembly because, you know, why not make it really zippy and let it run right in your browser?

Warning: You may find yourself endlessly comparing things like your grocery list to a random Shakespeare sonnet just because it’s that fun...

## How to Build and Deploy

To build and deploy the project, follow these steps:

1. **Clone the repository:**
    ```sh
    git clone https://github.com/achu-1612/fmt.git
    cd fmt
    ```

2. **Build the WebAssembly binary:**
    ```sh
    make build
    ```

3. **Build the Docker image:**
    ```sh
    make docker
    ```

4. **Run the Docker container:**
    ```sh
    make run
    ```

The application will be available at `http://localhost:3000`.

## How to Use the Application (or, "How to Waste an Afternoon Finding Every Typo")

1. Open the application – In your web browser, not in your toaster, and preferably before you get too emotionally attached to either piece of text.
2. Type your masterpiece – Drop your first text in the left box. Think of it as your magnum opus, even if it’s just a grocery list.
3. Type in the second piece – Enter the second text in the right box. Maybe it's a revised version, or maybe you’re just curious how Shakespeare stacks up against your shopping list.
5. Pick a battle style – Choose your comparison mode. Are you ready to go Character by Character or brave the intensity of a Word-by-Word face-off?
6. Press “Compare” – Brace yourself as the app reveals every difference, typo, and sneaky “the the” you somehow missed.
7. Click “Clear” or “Swap” – The "Clear" button erases everything for a fresh start, and the "Swap" button... you guessed it, swaps the texts. For those moments when you just have to see it from the other side.

## Additional (Nerdy) Details

- The WebAssembly binary? Oh yeah, it’s built with the Go compiler (using GOOS=js and GOARCH=wasm), because why keep things simple?
- The app is served by an Apache HTTP server via a Dockerfile, making this tool only slightly easier to run than a spaceship.
- Dark Mode – Because, of course, we have a dark mode. It’s 2024; you deserve to compare texts without straining your eyes.

Need more details? Check out the code and comments (non existent) in the repo... if you’re up for it.
