<script>
    import Item from "./Item.svelte"
    import { onMount} from "svelte"

    let scissors = []
    onMount(async () => {
        const res = await fetch("http://localhost:3000/scissor")
        scissors = await res.json()
    })
</script>

{#each scissors as scissor}
    <Item scissor={scissor} />
{/each}