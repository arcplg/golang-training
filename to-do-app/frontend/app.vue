<template>
  <div>
    <NuxtLayout>
      <NuxtPage/>
    </NuxtLayout>
  </div>
</template>

<script>
import LoginForm from '~/components/LoginForm.vue';
import Header from '~/components/layouts/Header.vue';

export default {
  components: {
    LoginForm,
    Header
  }
};
</script>
