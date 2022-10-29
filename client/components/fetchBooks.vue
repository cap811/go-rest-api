<template>
  <div class="container mx-auto">
    <h1 class="font-bold text-3xl pt-2">Books</h1>
      <ol>
        <li v-for="book in books" :key="book.id" class="pt-3">
          <h1 class="text-2xl font-semibold">
            {{ book.title }}
          </h1>
          <p>{{ book.author }}</p>
          <p class="italic">{{ book.year }}</p>
        </li>
      </ol>
  </div>
</template>

<script>
export default {
  name : "fetchBooks.vue",
  data() {
    return {
      books: []
    }
  },
  async fetch() {
    this.books = await fetch(
      'http://localhost/api/books'
    ).then(res => res.json())
  }
}
</script>

