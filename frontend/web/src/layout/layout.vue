<template>
	<router-link to="/">Home</router-link> |
	<router-link to="/index">Diagram</router-link> |
	<router-link to="/panda">Visual Script</router-link> |
	<router-link to="/http404">Budget</router-link>
	<router-view></router-view>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
	name: "Layout",
	components: {},
	computed: {},
	methods: {},
	watch: {},
	onCreated: () => {},
	onMounted: () => {},
});
</script>

<style></style>
