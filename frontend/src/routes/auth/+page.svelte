<svelte:head>
	<title>SpotDrop</title>
</svelte:head>


<script lang="ts">
 import * as Tabs from "$lib/components/ui/tabs/index";
 import * as Card from "$lib/components/ui/card/index";
 import { Button } from "$lib/components/ui/button/index";
 import { Input } from "$lib/components/ui/input/index";
 import { Label } from "$lib/components/ui/label/index";
</script>
 
<div class="flex h-[75%] w-full items-center justify-center bg-background">

    <div class="-mb-4 flex w-full max-w-sm flex-col gap-6">
    <Tabs.Root value="login">
    <Tabs.List>
    <Tabs.Trigger value="login">Login</Tabs.Trigger>
    <Tabs.Trigger value="signup">Sign Up</Tabs.Trigger>
    </Tabs.List>
    <Tabs.Content value="login">
    <Card.Root>
        <Card.Header>
        <Card.Title>Login</Card.Title>
        <Card.Description>
        Login to your account. If you don&apos;t have an account, you can create one by clicking the Sign Up tab.
        </Card.Description>
        </Card.Header>
        <Card.Content class="grid gap-6">
        <div class="grid gap-3">
        <Label for="tabs-username">Username</Label>
        <Input id="tabs-username" placeholder="Username" />
        </div>
        <div class="grid gap-3">
        <Label for="tabs-password">Password</Label>
        <Input id="tabs-password" placeholder="Password" />
        </div>
        </Card.Content>
        <Card.Footer>
        <Button>Login</Button>
        </Card.Footer>
    </Card.Root>
    </Tabs.Content>
    <Tabs.Content value="signup">
    <Card.Root>
        <Card.Header>
        <Card.Title>Sign Up</Card.Title>
        <Card.Description>
        Create a new account. If you already have an account, you can login by clicking the Login tab.
        </Card.Description>
        </Card.Header>
        <Card.Content class="grid gap-6">
        <div class="grid gap-3">
        <Label for="tabs-signup-username">Username</Label>
        <Input id="tabs-signup-username" placeholder="Username" />
        </div>
        <div class="grid gap-3">
        <Label for="tabs-signup-password">Password</Label>
        <Input id="tabs-signup-password" placeholder="Password" />
        </div>
        </Card.Content>
        <Card.Footer>
        <Button>Sign Up</Button>
        </Card.Footer>
    </Card.Root>
    </Tabs.Content>
    </Tabs.Root>
    </div>
</div>
