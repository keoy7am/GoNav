<template>
    <div class="min-h-screen flex items-center justify-center bg-gray-50 dark:bg-gray-900 py-12 px-4 sm:px-6 lg:px-8">
        <div class="max-w-sm w-full space-y-8">
            <div>
                <BookmarkIcon class="mx-auto h-12 w-auto text-blue-500 dark:text-blue-400" />
                <h2 class="mt-6 text-center text-3xl font-extrabold text-gray-900 dark:text-white">
                    {{ t('register.title') }}
                </h2>
            </div>
            <form class="mt-8 space-y-6" @submit.prevent="register">
                <div class="rounded-md shadow-sm -space-y-px">
                    <div>
                        <label for="username" class="sr-only">{{ t('register.username') }}</label>
                        <input id="username" name="username" type="text" required
                               class="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 dark:border-gray-700 placeholder-gray-500 dark:placeholder-gray-400 text-gray-900 dark:text-white rounded-t-md focus:outline-none focus:ring-blue-500 focus:border-blue-500 dark:focus:ring-blue-400 dark:focus:border-blue-400 focus:z-10 sm:text-sm bg-white dark:bg-gray-800"
                               :placeholder="t('register.username')" v-model="username">
                    </div>
                    <div>
                        <label for="email" class="sr-only">{{ t('register.email') }}</label>
                        <input id="email" name="email" type="email" autocomplete="email" required
                               class="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 dark:border-gray-700 placeholder-gray-500 dark:placeholder-gray-400 text-gray-900 dark:text-white focus:outline-none focus:ring-blue-500 focus:border-blue-500 dark:focus:ring-blue-400 dark:focus:border-blue-400 focus:z-10 sm:text-sm bg-white dark:bg-gray-800"
                               :placeholder="t('register.email')" v-model="email">
                    </div>
                    <div>
                        <label for="password" class="sr-only">{{ t('register.password') }}</label>
                        <input id="password" name="password" type="password" autocomplete="new-password" required
                               class="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 dark:border-gray-700 placeholder-gray-500 dark:placeholder-gray-400 text-gray-900 dark:text-white rounded-b-md focus:outline-none focus:ring-blue-500 focus:border-blue-500 dark:focus:ring-blue-400 dark:focus:border-blue-400 focus:z-10 sm:text-sm bg-white dark:bg-gray-800"
                               :placeholder="t('register.password')" v-model="password">
                    </div>
                </div>

                <div>
                    <button type="submit"
                            class="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500">
                        {{ t('register.submit') }}
                    </button>
                </div>
            </form>
            <div class="text-center">
                <router-link to="/login" class="font-medium text-blue-600 hover:text-blue-500 dark:text-blue-400 dark:hover:text-blue-300">
                    {{ t('register.loginLink') }}
                </router-link>
            </div>
        </div>
    </div>
</template>

<script>
import { ref } from 'vue'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import BookmarkIcon from '@/components/BookmarkIcon.vue'
import { useToast } from "vue-toastification"

export default {
    name: 'Register',
    components: {
        BookmarkIcon
    },
    setup() {
        const store = useStore()
        const router = useRouter()
        const { t } = useI18n()
        const username = ref('')
        const email = ref('')
        const password = ref('')
        const toast = useToast()

        const register = async () => {
            try {
                await store.dispatch('register', {
                    username: username.value,
                    email: email.value,
                    password: password.value
                })
                router.push('/login')
                toast.success(t('register.success'))
            } catch (error) {
                console.error('Registration error:', error)
                toast.error(t('register.error'))
            }
        }

        return {
            username,
            email,
            password,
            register,
            t
        }
    }
}
</script>