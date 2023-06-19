<script>
    import Card from "./Card.svelte"
    import Modal from "./Modal.svelte"
    import { Modals, closeModal, openModal } from "svelte-modals"

    export let scissor
    let showCard = true
    async function update(data) {
        const json = {
            redirect: data.redirect,
            scissor: data.scissor,
            random: data.random,
            id: scissor.id
        }

        await fetch("http://localhost:3000/scissor", {
            method: "PATCH",
            headers: {"Content-Type": "application/json"},
            body: JSON.stringify(json)
        }).then(res => {
            console.log(res)
        })
    }

    function handleOpen(scissor) {
        openModal(Modal, {
            title: "Update Scissor Link",
            send: update,
            scissor: scissor.scissor,
            redirect: scissor.redirect,
            random: scissor.random
        })
    }

    async function deleteScissor() {
        if (confirm("Are you sure you wish to delete this Scissor link (" + scissor.scissor + ")?")) {
            await fetch("http://localhost:3000/scissor/" + scissor.id, {
                method: "DELETE"
            }).then(response => {
                showCard = false
                console.log(response)
            })
        }
    }

</script>

{#if showCard }
<Card>
    <p>Scissor: http://localhost:3000/r/{scissor.scissor} </p>
    <p>Redirect: {scissor.redirect}</p>
    <p>Clicked: {scissor.clicked}</p>
    <button class="update" on:click={ handleOpen(scissor) }>Update</button>
    <button class="delete" on:click={deleteScissor}>Delete</button>
</Card>
{/if}
<Modals>
    <div 
        slot="backdrop"
        class="backdrop"
        on:click={closeModal}
    />
</Modals>


<style>
    button{
        color: white;
        font-weight: bolder;
        border: none;
        padding: .75rem;
        border-radius: 4px;
    }
    .update{
        background-color: yellowgreen;
    }
    .delete{
        background-color: red;
    }
    .backdrop {
        position: fixed;
        top: 0;
        bottom: 0;
        right: 0;
        left: 0;
        background: rgb(255, 255, 255)
    }
</style>