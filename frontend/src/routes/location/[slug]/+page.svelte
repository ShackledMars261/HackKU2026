<script>
  import { onMount } from 'svelte';

  /**
   * Props shaped after the domain types:
   *
   * Location = { id, location: GeoJSON, owner, name, overallRating }
   * Post     = { id, userId, locationId, rating, description, createdAt, photoFileIds }
   *
   * Usage:
   *   <PlaceDetail data={{ location, posts }} />
   */

  let {
    data = {
      location: /** @type {import('./types').Location} */ ({
        id: 'loc-1',
        name: 'Trattoria dal Vecchio',
        owner: 'user-42',
        overallRating: 4,
        location: {
          type: 'Feature',
          geometry: { type: 'Point', coordinates: [11.2558, 43.7696] },
          properties: { address: '14 Via Roma, Florence, Italy' },
        },
      }),
      posts: /** @type {import('./types').Post[]} */ ([
        { id: 'p1', userId: 'u1', locationId: 'loc-1', rating: 5, description: 'Absolutely incredible pasta. The ambience is perfect for a date night.', createdAt: '2024-11-03T18:22:00Z', photoFileIds: ['https://images.unsplash.com/photo-1555396273-367ea4eb4db5?w=800&q=80'] },
        { id: 'p2', userId: 'u2', locationId: 'loc-1', rating: 4, description: 'Fantastic service and great wine list. Would definitely return.',          createdAt: '2024-10-21T20:05:00Z', photoFileIds: ['https://images.unsplash.com/photo-1414235077428-338989a2e8c0?w=800&q=80'] },
        { id: 'p3', userId: 'u3', locationId: 'loc-1', rating: 4, description: 'Lovely spot, a bit noisy on weekends but the food makes up for it.',        createdAt: '2024-09-14T13:44:00Z', photoFileIds: ['https://images.unsplash.com/photo-1504674900247-0877df9cc836?w=800&q=80'] },
        { id: 'p4', userId: 'u4', locationId: 'loc-1', rating: 3, description: 'Good traditional food, nothing too surprising but solid quality.',           createdAt: '2024-08-30T12:10:00Z', photoFileIds: ['https://images.unsplash.com/photo-1482049016688-2d3e1b311543?w=800&q=80'] },
        { id: 'p5', userId: 'u5', locationId: 'loc-1', rating: 5, description: 'The tiramisu is the best I have ever had. Truly unforgettable evening.',    createdAt: '2024-07-19T19:55:00Z', photoFileIds: [] },
      ]),
    },
  } = $props();

  // ── Derived from Location ──────────────────────────────────────────────────
  /** Human-readable address stored in GeoJSON feature properties */
  let address = $derived(
    data.location.location?.properties?.address ?? ''
  );

  /** All photo URLs collected across all posts — used for the carousel */
  let carouselImages = $derived(
    data.posts.flatMap(p => p.photoFileIds).filter(Boolean)
  );

  // ── Carousel state ─────────────────────────────────────────────────────────
  let currentImage = $state(0);
  let isTransitioning = $state(false);

  function goTo(idx) {
    if (isTransitioning || carouselImages.length === 0) return;
    isTransitioning = true;
    currentImage = ((idx % carouselImages.length) + carouselImages.length) % carouselImages.length;
    setTimeout(() => isTransitioning = false, 300);
  }

  // ── Reviews / posts state ──────────────────────────────────────────────────
  let posts = $state(data.posts);
  let visibleCount = $state(4);

  // ── Add-review modal ───────────────────────────────────────────────────────
  let showModal = $state(false);
  let newPost = $state({ description: '', rating: 5 });

  function submitReview() {
    if (!newPost.description.trim()) return;
    /** @type {import('./types').Post} */
    const post = {
      id: crypto.randomUUID(),
      userId: 'current-user',          // replace with real session.userId
      locationId: data.location.id,
      rating: newPost.rating,
      description: newPost.description,
      createdAt: new Date().toISOString(),
      photoFileIds: [],
    };
    posts = [post, ...posts];
    newPost = { description: '', rating: 5 };
    showModal = false;
  }

  // ── Helpers ────────────────────────────────────────────────────────────────
  function stars(n, max = 5) {
    return Array.from({ length: max }, (_, i) => i < n ? '★' : '☆').join('');
  }

  /** Deterministic avatar placeholder seeded from userId */
  function avatarUrl(userId) {
    const seed = userId.split('').reduce((a, c) => a + c.charCodeAt(0), 0) % 70;
    return `https://i.pravatar.cc/64?img=${seed}`;
  }

  /** e.g. "3 Oct 2024" */
  function fmtDate(iso) {
    return new Date(iso).toLocaleDateString('en-GB', { day: 'numeric', month: 'short', year: 'numeric' });
  }

  // ── Keyboard nav ──────────────────────────────────────────────────────────
  function handleKey(e) {
    if (e.key === 'ArrowLeft')  goTo(currentImage - 1);
    if (e.key === 'ArrowRight') goTo(currentImage + 1);
    if (e.key === 'Escape') showModal = false;
  }

  onMount(() => {
    window.addEventListener('keydown', handleKey);
    return () => window.removeEventListener('keydown', handleKey);
  });
</script>

<!-- ─────────────────────────────────────────────────────────────────────────
     Root layout: two-panel
──────────────────────────────────────────────────────────────────────────── -->
<div class="flex h-full w-full gap-4 p-4 overflow-hidden font-sans">

  <!-- ══ LEFT PANEL: Location detail card ══════════════════════════════════ -->
  <div
    class="flex flex-col overflow-hidden flex-shrink-0 border border-border bg-card shadow-xl"
    style="width: 340px; border-radius: 28px;"
  >

    <!-- Carousel (photos collected from post.photoFileIds) -->
    <div class="relative flex-shrink-0 overflow-hidden bg-muted" style="height: 260px; border-radius: 20px 20px 0 0;">
      {#if carouselImages.length > 0}
        {#each carouselImages as src, i}
          <img
            {src}
            alt="Location photo {i + 1}"
            class="absolute inset-0 w-full h-full object-cover transition-opacity duration-300"
            style="opacity: {i === currentImage ? 1 : 0}; pointer-events: {i === currentImage ? 'auto' : 'none'};"
          />
        {/each}
      {:else}
        <div class="absolute inset-0 flex items-center justify-center text-muted-foreground text-[13px]">
          No photos yet
        </div>
      {/if}

      <!-- Gradient -->
      <div class="absolute inset-0" style="background: linear-gradient(to top, rgba(0,0,0,.45) 0%, transparent 55%);"></div>

      <!-- Controls (only when multiple images) -->
      {#if carouselImages.length > 1}
        <div class="absolute bottom-3 left-0 right-0 flex items-center justify-center gap-2">
          <button onclick={() => goTo(currentImage - 1)} class="text-white/80 hover:text-white transition-colors text-lg leading-none" aria-label="Previous">‹</button>
          {#each carouselImages as _, i}
            <button
              onclick={() => goTo(i)}
              class="rounded-full transition-all duration-200"
              style="width: {i === currentImage ? '18px' : '7px'}; height: 7px; background: {i === currentImage ? 'white' : 'rgba(255,255,255,0.45)'};"
              aria-label="Image {i + 1}"
            ></button>
          {/each}
          <button onclick={() => goTo(currentImage + 1)} class="text-white/80 hover:text-white transition-colors text-lg leading-none" aria-label="Next">›</button>
        </div>
      {/if}
    </div>

    <!-- Location info (from data.location) -->
    <div class="flex flex-col gap-3 px-6 py-5 flex-1 overflow-y-auto">
      <!-- location.name -->
      <h2 class="text-xl font-bold tracking-tight text-foreground leading-snug">
        {data.location.name}
      </h2>

      <div class="flex flex-col gap-1.5 text-[13px]">
        <!-- GeoJSON Feature properties.address -->
        {#if address}
          <div class="flex gap-1.5">
            <span class="font-semibold text-foreground">Address:</span>
            <span class="text-muted-foreground">{address}</span>
          </div>
        {/if}

        <!-- location.overallRating -->
        <div class="flex items-center gap-1.5">
          <span class="font-semibold text-foreground">Avg. Rating:</span>
          <span class="text-amber-500 tracking-tight">{stars(data.location.overallRating)}</span>
          <span class="text-muted-foreground text-[12px]">{data.location.overallRating}/5</span>
        </div>

        <!-- Derived post count -->
        <div class="flex items-center gap-1.5">
          <span class="font-semibold text-foreground">Reviews:</span>
          <span class="text-muted-foreground">{posts.length}</span>
        </div>
      </div>

      <!-- Rating fill bar -->
      <div class="mt-1">
        <div class="h-2 w-full rounded-full overflow-hidden" style="background: var(--muted);">
          <div
            class="h-full rounded-full transition-all duration-700"
            style="width: {(data.location.overallRating / 5) * 100}%; background: var(--primary); opacity: 0.75;"
          ></div>
        </div>
        <p class="text-[10px] text-muted-foreground mt-1">Overall rating</p>
      </div>
    </div>
  </div>

  <!-- ══ RIGHT PANEL: Posts (reviews) ══════════════════════════════════════ -->
  <div
    class="flex flex-col flex-1 overflow-hidden border border-border bg-card shadow-xl"
    style="border-radius: 28px;"
  >
    <!-- Header -->
    <div class="flex-shrink-0 px-6 pt-5 pb-3 border-b border-border">
      <h3 class="text-[11px] font-semibold uppercase tracking-widest text-muted-foreground">Reviews</h3>
    </div>

    <!-- Post rows -->
    <div class="flex-1 overflow-y-auto px-4 py-2">
      {#each posts.slice(0, visibleCount) as post (post.id)}
        <div class="flex gap-3 px-2 py-3 border-b border-border/50 last:border-0 hover:bg-accent/10 transition-colors rounded-xl">
          <!-- Avatar seeded from post.userId -->
          <img
            src={avatarUrl(post.userId)}
            alt="User {post.userId}"
            class="w-10 h-10 rounded-full object-cover flex-shrink-0 ring-2 ring-border"
          />
          <div class="flex flex-col gap-0.5 min-w-0">
            <div class="flex items-center gap-2 flex-wrap">
              <!-- post.userId (swap for resolved username when available) -->
              <span class="text-[13px] font-semibold text-foreground leading-none truncate max-w-[120px]">
                {post.userId}
              </span>
              <span class="text-amber-500 text-[11px] tracking-tight">{stars(post.rating)}</span>
              <!-- post.createdAt -->
              <span class="text-[11px] text-muted-foreground ml-auto">{fmtDate(post.createdAt)}</span>
            </div>
            <!-- post.description -->
            <p class="text-[12px] text-muted-foreground leading-relaxed line-clamp-2">{post.description}</p>
            <!-- post.photoFileIds thumbnail strip -->
            {#if post.photoFileIds.length > 0}
              <div class="flex gap-1 mt-1.5">
                {#each post.photoFileIds.slice(0, 3) as src}
                  <img {src} alt="Post photo" class="w-10 h-10 rounded-lg object-cover ring-1 ring-border" />
                {/each}
              </div>
            {/if}
          </div>
        </div>
      {/each}

      {#if posts.length === 0}
        <p class="text-center text-[13px] text-muted-foreground py-10">No reviews yet. Be the first!</p>
      {/if}
    </div>

    <!-- Footer -->
    <div class="flex-shrink-0 flex items-center justify-end gap-2 px-6 py-4 border-t border-border">
      {#if visibleCount < posts.length}
        <button
          onclick={() => visibleCount += 4}
          class="rounded-full border border-border bg-background px-5 py-2 text-[12px] font-semibold text-foreground hover:bg-accent hover:text-accent-foreground transition-colors"
        >
          View More
        </button>
      {/if}
      <button
        onclick={() => showModal = true}
        class="rounded-full px-5 py-2 text-[12px] font-semibold text-primary-foreground transition-all hover:opacity-90 active:scale-95"
        style="background: var(--primary);"
      >
        Add Review
      </button>
    </div>
  </div>

</div>

<!-- ══ Add Review Modal ═══════════════════════════════════════════════════════ -->
{#if showModal}
  <div
    class="fixed inset-0 z-50 flex items-center justify-center"
    style="background: rgba(0,0,0,0.4); backdrop-filter: blur(4px);"
    onclick={e => { if (e.target === e.currentTarget) showModal = false; }}
  >
    <div
      class="w-full max-w-sm mx-4 border border-border bg-card shadow-2xl flex flex-col gap-4 p-6"
      style="border-radius: 24px;"
    >
      <h3 class="text-base font-bold text-foreground tracking-tight">Add a Review</h3>

      <textarea
        placeholder="Share your experience…"
        bind:value={newPost.description}
        rows="4"
        class="w-full rounded-xl border border-border bg-background px-4 py-2.5 text-[13px] text-foreground placeholder:text-muted-foreground outline-none focus:ring-1 focus:ring-ring transition resize-none"
      ></textarea>

      <!-- Star picker -->
      <div class="flex items-center gap-1">
        <span class="text-[12px] text-muted-foreground mr-1">Rating:</span>
        {#each [1,2,3,4,5] as n}
          <button
            onclick={() => newPost.rating = n}
            class="text-xl transition-transform hover:scale-110"
            style="color: {n <= newPost.rating ? '#f59e0b' : 'var(--muted-foreground)'};"
            aria-label="Rate {n} stars"
          >★</button>
        {/each}
      </div>

      <div class="flex gap-2 justify-end">
        <button
          onclick={() => showModal = false}
          class="rounded-full border border-border bg-background px-5 py-2 text-[12px] font-semibold text-foreground hover:bg-accent transition-colors"
        >Cancel</button>
        <button
          onclick={submitReview}
          class="rounded-full px-5 py-2 text-[12px] font-semibold text-primary-foreground hover:opacity-90 active:scale-95 transition-all"
          style="background: var(--primary);"
        >Submit</button>
      </div>
    </div>
  </div>
{/if}