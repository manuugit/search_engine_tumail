<script>
import SearchTab from './SearchTab.vue'
import MailsList from './MailsList.vue'
import { getMailsxMatch } from './../services/connectBack'

export default {
    methods: {
        async searchMatch(word) {
            this.searchedWord = word;
            this.mails = await this.getMails()
        },
        async getMails() {
            const term = this.searchedWord
            // timestamps del momento en que se cargó el registro en zincsearch
            const startTime = '2023-08-02T14:28:31.894Z'
            const endTime = '2023-08-04T15:28:31.894Z'
            const maxResults = 50
            const fields = []
            const data = await getMailsxMatch(term, startTime, endTime, maxResults, fields)
            return data.hits
        }
    },
    data() {
        return {
            searchedWord: '',
            mails: [] 
        }
    },
    components: {
        SearchTab,
        MailsList
    },
}
</script>

<template>
    <SearchTab @send-word="searchMatch" />
    <MailsList :mails="mails" />
</template>