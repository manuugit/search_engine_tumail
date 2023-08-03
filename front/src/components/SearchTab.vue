<script>
export default {
  // reactive state
  data() {
    return {
      word: ''
    }
  },

  // functions that mutate state and trigger updates
  methods: {
    sendWord() {
      // se crea un evento para obtener el valor de word en un componente externo
      // pasamos word en el evento
      this.$emit('send-word', this.word);
    }
  },
  // lifecycle hooks

}
</script>

<template>
  <div class="flex flex-row margins-container">
    <img class="h-12 w-12 flex-none rounded-full bg-gray-50 search-image"
      src='https://w7.pngwing.com/pngs/385/718/png-transparent-computer-icons-encapsulated-postscript-search-box-search-autocomplete-symbol-search-box.png'
      alt="" />
    <input class="mid-3-4-space" v-model="word" placeholder="Word" autofocus />
    <button class="button-style" @click="sendWord">Search</button>
  </div>
</template>
