package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", serveHomePage)
	http.HandleFunc("/about", serveAboutPage)
	http.HandleFunc("/contact", serveContactPage)
	http.HandleFunc("/projects", serveProjectsPage)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func serveHomePage(w http.ResponseWriter, r *http.Request) {
	html := `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Home Page</title>
    <link href="https://cdn.jsdelivr.net/npm/tailwindcss@2.2.19/dist/tailwind.min.css" rel="stylesheet">
</head>
<body class="bg-gray-100 text-gray-900">

    <header class="bg-blue-500 text-white p-4">
        <nav class="container mx-auto flex items-center justify-between">
            <a href="/" class="text-2xl font-bold">My Website</a>
            <div>
                <a href="/" class="text-white hover:underline px-4">Home</a>
                <a href="/about" class="text-white hover:underline px-4">About</a>
                <a href="/contact" class="text-white hover:underline px-4">Contact</a>
                <a href="/projects" class="text-white hover:underline px-4">Projects</a>
            </div>
        </nav>
    </header>

    <main class="container mx-auto p-4 bg-white rounded-lg shadow-md mb-8">
        <section class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-3">
            <div class="bg-gray-50 p-4 rounded-lg shadow-md flex flex-col items-center">
                <img src="/img/gallery1.jpg" alt="Sample Image" class="w-32 h-32 mb-4 rounded-full">
                <h2 class="text-xl font-semibold mb-2">Item 1</h2>
                <p class="text-center">This is a description of item 1. It has some details about the content.</p>
            </div>
            <div class="bg-gray-50 p-6 rounded-lg shadow-md flex flex-col items-center">
                <img src="/img/gallery1.jpg" alt="Sample Image" class="w-32 h-32 mb-4 rounded-full">
                <h2 class="text-xl font-semibold mb-2">Item 2</h2>
                <p class="text-center">This is a description of item 2. It has some details about the content.</p>
            </div>
            <div class="bg-gray-50 p-6 rounded-lg shadow-md flex flex-col items-center">
                <img src="/img/gallery1.jpg" alt="Sample Image" class="w-32 h-32 mb-4 rounded-full">
                <h2 class="text-xl font-semibold mb-2">Item 3</h2>
                <p class="text-center">This is a description of item 3. It has some details about the content.</p>
            </div>
        </section>
    </main>

    <footer class="bg-blue-500 text-white p-4 text-center rounded-lg shadow-md mb-8">
        <p>&copy; 2024 My Website</p>
    </footer>

</body>
</html>
`
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}

func serveAboutPage(w http.ResponseWriter, r *http.Request) {
	html := `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>About Us</title>
    <link href="https://cdn.jsdelivr.net/npm/tailwindcss@2.2.19/dist/tailwind.min.css" rel="stylesheet">
</head>
<body class="bg-gray-100 text-gray-900">

    <header class="bg-blue-500 text-white p-4">
        <nav class="container mx-auto flex items-center justify-between">
            <a href="/" class="text-2xl font-bold">My Website</a>
            <div>
                <a href="/" class="text-white hover:underline px-4">Home</a>
                <a href="/about" class="text-white hover:underline px-4">About</a>
                <a href="/contact" class="text-white hover:underline px-4">Contact</a>
                <a href="/projects" class="text-white hover:underline px-4">Projects</a>
            </div>
        </nav>
    </header>

    <main class="container mx-auto p-8 bg-white rounded-lg shadow-md mb-8">
        <section class="grid grid-cols-1 md:grid-cols-2 gap-8">
            <div class="bg-gray-50 p-6 rounded-lg shadow-md">
                <h2 class="text-2xl font-semibold mb-4">Our Mission</h2>
                <p>Our mission is to provide high-quality content and services that exceed our clients' expectations.</p>
            </div>
            <div class="bg-gray-50 p-6 rounded-lg shadow-md">
                <h2 class="text-2xl font-semibold mb-4">Our Team</h2>
                <p>Meet our team of experts who are dedicated to delivering exceptional results.</p>
            </div>
        </section>
    </main>

    <footer class="bg-blue-500 text-white p-4 text-center rounded-lg shadow-md mb-8">
        <p>&copy; 2024 My Website</p>
    </footer>

</body>
</html>
`
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}

func serveContactPage(w http.ResponseWriter, r *http.Request) {
	html := `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Contact Us</title>
    <link href="https://cdn.jsdelivr.net/npm/tailwindcss@2.2.19/dist/tailwind.min.css" rel="stylesheet">
</head>
<body class="bg-gray-100 text-gray-900">

    <header class="bg-blue-500 text-white p-4">
        <nav class="container mx-auto flex items-center justify-between">
            <a href="/" class="text-2xl font-bold">My Website</a>
            <div>
                <a href="/" class="text-white hover:underline px-4">Home</a>
                <a href="/about" class="text-white hover:underline px-4">About</a>
                <a href="/contact" class="text-white hover:underline px-4">Contact</a>
                <a href="/projects" class="text-white hover:underline px-4">Projects</a>
            </div>
        </nav>
    </header>

    <main class="container mx-auto p-8 bg-white rounded-lg shadow-md mb-8">
        <section class="grid grid-cols-1 md:grid-cols-2 gap-8">
            <div class="bg-gray-50 p-6 rounded-lg shadow-md">
                <h2 class="text-2xl font-semibold mb-4">Get in Touch</h2>
                <p>If you have any questions or inquiries, please reach out to us via the contact form or email.</p>
            </div>
            <div class="bg-gray-50 p-6 rounded-lg shadow-md">
                <h2 class="text-2xl font-semibold mb-4">Contact Form</h2>
                <form>
                    <div class="mb-4">
                        <label for="name" class="block text-sm font-medium mb-2">Name</label>
                        <input type="text" id="name" class="w-full p-2 border rounded">
                    </div>
                    <div class="mb-4">
                        <label for="email" class="block text-sm font-medium mb-2">Email</label>
                        <input type="email" id="email" class="w-full p-2 border rounded">
                    </div>
                    <div class="mb-4">
                        <label for="message" class="block text-sm font-medium mb-2">Message</label>
                        <textarea id="message" rows="4" class="w-full p-2 border rounded"></textarea>
                    </div>
                    <button type="submit" class="bg-blue-500 text-white p-2 rounded">Send</button>
                </form>
            </div>
        </section>
    </main>

    <footer class="bg-blue-500 text-white p-4 text-center rounded-lg shadow-md mb-8">
        <p>&copy; 2024 My Website</p>
    </footer>

</body>
</html>
`
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}

func serveProjectsPage(w http.ResponseWriter, r *http.Request) {
	html := `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Projects</title>
    <link href="https://cdn.jsdelivr.net/npm/tailwindcss@2.2.19/dist/tailwind.min.css" rel="stylesheet">
</head>
<body class="bg-gray-100 text-gray-900">

    <header class="bg-blue-500 text-white p-4">
        <nav class="container mx-auto flex items-center justify-between">
            <a href="/" class="text-2xl font-bold">My Website</a>
            <div>
                <a href="/" class="text-white hover:underline px-4">Home</a>
                <a href="/about" class="text-white hover:underline px-4">About</a>
                <a href="/contact" class="text-white hover:underline px-4">Contact</a>
                <a href="/projects" class="text-white hover:underline px-4">Projects</a>
            </div>
        </nav>
    </header>

    <main class="container mx-auto p-8 bg-white rounded-lg shadow-md mb-8">
        <section class="grid grid-cols-1 md:grid-cols-2 gap-8">
            <div class="bg-gray-50 p-6 rounded-lg shadow-md">
                <h2 class="text-2xl font-semibold mb-4">Get in Touch</h2>
                <p>If you have any questions or inquiries, please reach out to us via the contact form or email.</p>
            </div>
            <div class="bg-gray-50 p-6 rounded-lg shadow-md">
                <h2 class="text-2xl font-semibold mb-4">Contact Form</h2>
                <form>
                    <div class="mb-4">
                        <label for="name" class="block text-sm font-medium mb-2">Name</label>
                        <input type="text" id="name" class="w-full p-2 border rounded">
                    </div>
                    <div class="mb-4">
                        <label for="email" class="block text-sm font-medium mb-2">Email</label>
                        <input type="email" id="email" class="w-full p-2 border rounded">
                    </div>
                    <div class="mb-4">
                        <label for="message" class="block text-sm font-medium mb-2">Message</label>
                        <textarea id="message" rows="4" class="w-full p-2 border rounded"></textarea>
                    </div>
                    <button type="submit" class="bg-blue-500 text-white p-2 rounded">Send</button>
                </form>
            </div>
        </section>
    </main>

    <footer class="bg-blue-500 text-white p-4 text-center rounded-lg shadow-md mb-8">
        <p>&copy; 2024 My Website</p>
    </footer>

</body>
</html>
`
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}
