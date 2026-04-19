<script lang="ts">
	import * as Dialog from '$lib/components/ui/dialog/index';
	import { Button } from '$lib/components/ui/button/index';

	let open = $state(false);
	let file = $state<File | null>(null);
	let loading = $state(false);
	let error = $state<string | null>(null);
	let success = $state(false);
	let dragover = $state(false);

	function onFileChange(e: Event) {
		const input = e.currentTarget as HTMLInputElement;
		file = input.files?.[0] ?? null;
		error = null;
		success = false;
	}

	function onDrop(e: DragEvent) {
		e.preventDefault();
		dragover = false;
		file = e.dataTransfer?.files?.[0] ?? null;
		error = null;
		success = false;
	}

	async function handleUpload() {
		if (!file) {
			error = 'Please select a file.';
			return;
		}

		loading = true;
		error = null;
		success = false;

		try {
			const resp = await fetch('/api/asset', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/octet-stream',
					'X-Filename': file.name
				},
				body: file
			});

			if (!resp.ok) throw new Error(await resp.text());

			success = true;
			file = null;
		} catch (e: any) {
			error = e.message;
		} finally {
			loading = false;
		}
	}

	function onClose() {
		if (loading) return;
		open = false;
		file = null;
		error = null;
		success = false;
	}
</script>

<Dialog.Root bind:open>
	<Dialog.Trigger>
		{#snippet child({ props })}
			<button
				{...props}
				class="inline-flex items-center justify-center gap-1 rounded-md px-4 py-2 text-sm font-medium transition-colors hover:bg-accent hover:text-accent-foreground focus:outline-none"
			>
				Upload Image
			</button>
		{/snippet}
	</Dialog.Trigger>

	<Dialog.Content class="sm:max-w-md">
		<Dialog.Header>
			<Dialog.Title>Upload Image</Dialog.Title>
			<Dialog.Description>Select or drag and drop an image to upload.</Dialog.Description>
		</Dialog.Header>

		<div class="flex flex-col gap-4 py-2">
			<!-- Drop zone -->
			<label
				class="flex cursor-pointer flex-col items-center justify-center gap-3 rounded-xl border-2 border-dashed p-8 transition-colors
                    {dragover
					? 'border-primary bg-primary/5'
					: 'border-border hover:border-primary/50 hover:bg-muted/50'}"
				ondragover={(e) => {
					e.preventDefault();
					dragover = true;
				}}
				ondragleave={() => (dragover = false)}
				ondrop={onDrop}
			>
				<svg
					class="h-10 w-10 text-muted-foreground"
					xmlns="http://www.w3.org/2000/svg"
					fill="none"
					viewBox="0 0 24 24"
					stroke="currentColor"
				>
					<path
						stroke-linecap="round"
						stroke-linejoin="round"
						stroke-width="1.5"
						d="M3 16.5v2.25A2.25 2.25 0 005.25 21h13.5A2.25 2.25 0 0021 18.75V16.5m-13.5-9L12 3m0 0l4.5 4.5M12 3v13.5"
					/>
				</svg>

				{#if file}
					<div class="flex items-center gap-2 text-sm font-medium text-foreground">
						<svg class="h-4 w-4 text-primary" viewBox="0 0 20 20" fill="currentColor">
							<path
								fill-rule="evenodd"
								d="M4 4a2 2 0 012-2h4.586A2 2 0 0112 2.586L15.414 6A2 2 0 0116 7.414V16a2 2 0 01-2 2H6a2 2 0 01-2-2V4z"
								clip-rule="evenodd"
							/>
						</svg>
						{file.name}
						<span class="text-muted-foreground">({(file.size / 1024).toFixed(1)} KB)</span>
					</div>
				{:else}
					<div class="text-center">
						<p class="text-sm font-medium text-foreground">Drop a file here or click to browse</p>
						<p class="mt-1 text-xs text-muted-foreground">PNG, JPG, GIF, WEBP supported</p>
					</div>
				{/if}

				<input type="file" accept="image/*" class="hidden" onchange={onFileChange} />
			</label>

			{#if success}
				<p class="flex items-center gap-1.5 text-sm text-green-600">
					<svg class="h-4 w-4" viewBox="0 0 20 20" fill="currentColor">
						<path
							fill-rule="evenodd"
							d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"
							clip-rule="evenodd"
						/>
					</svg>
					File uploaded successfully!
				</p>
			{/if}

			{#if error}
				<p class="text-sm text-destructive">{error}</p>
			{/if}
		</div>

		<Dialog.Footer>
			<Button variant="secondary" onclick={onClose} disabled={loading}>Cancel</Button>
			<Button onclick={handleUpload} disabled={loading || !file}>
				{loading ? 'Uploading…' : 'Upload'}
			</Button>
		</Dialog.Footer>
	</Dialog.Content>
</Dialog.Root>
