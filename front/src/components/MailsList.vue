<script>
export default {
    props: {
        mails: {
            type: Array,
            required: true
        },
        searchedWord: {
            type: String,
            required: true
        }
    },
    methods: {
        selectMail(mail) {
            this.selectedMail = mail; // Asigna el valor del elemento seleccionado a la variable
            //console.log(this.selectedMail._source.Body)
        }
    },
    computed: {
        highlightContent() {
        const regex = new RegExp(`\\b${this.searchedWord}\\b`, 'gi');
        // contenido donde se van a buscar la palabra que queremos resaltar
        const content = this.selectedMail._source.Body;
            return content.replace(regex, '<span style="font-weight: bold;color: purple;">$&</span>');
        }
    },
    data() {
        return {
            selectedMail: null // Variable para almacenar el elemento seleccionado
        };
    },

}
</script>

<template>
    <div class="margins-container flex flex-row">
        <div class="col-start-1 col-span-4 mid-1-2-space">
            <p class="bodymail-style" v-if="mails.length"> **Click on the email to read the body** <br> </p>
            <ul role="list" class="divide-y divide-gray-100">
                <li v-for="mail in mails" :key="mail._source.MessageID" class="flex justify-between gap-x-6 py-5"
                    @click="selectMail(mail)" style="cursor: pointer;">
                    <div class="flex gap-x-4">
                        <img class="h-12 w-12 flex-none rounded-full bg-gray-50"
                            src='data:image/png;base64,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'
                            alt="" />
                        <div class="min-w-0 flex-auto">
                            <p class="text-sm font-semibold leading-6 text-gray-900"> {{ mail._source.Subject }}</p>
                            <p class="text-sm leading-6 text-gray-900"> {{ mail._source.CategoryFolder }}</p>
                            <p class="mt-1 truncate text-xs leading-5 text-gray-500"> From: {{ mail._source.SenderEmail }}</p>
                        </div>
                    </div>
                    <div class="hidden sm:flex sm:flex-col sm:items-end">
                        <p class="text-sm leading-6 text-gray-900"> .</p>
                        <p class="text-sm leading-6 text-gray-900"> {{ mail._source.Date }}</p>
                        <p class="mt-1 text-xs leading-5 text-gray-500"> To: {{ mail._source.RecipientEmail }}
                        </p>
                    </div>
                </li>
            </ul>
        </div>
        
        <span v-if="selectedMail != null" class="mail-container-style">
            <p class="bodymail-style"> <br>
                <span v-html="highlightContent"></span>
            </p>
        </span>
    </div>
</template>
  
