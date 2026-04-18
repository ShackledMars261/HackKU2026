<script>
  import { onMount, onDestroy } from 'svelte';

  let { data = [] } = $props();

  let mapEl;
  let map;
  let L;

  let filterText = $state('');
  let currentPage = $state(1);
  let sorting = $state({ column: null, direction: 'asc' });
  const pageSize = 8;

  const tableData = data.length ? data : [
    { id: 1,  city: 'Paris',     lat: 48.8584,  lng:  2.2945  },
    { id: 2,  city: 'Rome',      lat: 41.8902,  lng: 12.4922  },
    { id: 3,  city: 'Barcelona', lat: 41.4036,  lng:  2.1744  },
    { id: 4,  city: 'London',    lat: 51.5007,  lng: -0.1246  },
    { id: 5,  city: 'Athens',    lat: 37.9715,  lng: 23.7257  },
    { id: 6,  city: 'Amsterdam', lat: 52.3752,  lng:  4.8840  },
    { id: 7,  city: 'Berlin',    lat: 52.5163,  lng: 13.3777  },
    { id: 8,  city: 'Prague',    lat: 50.0865,  lng: 14.4114  },
    { id: 9,  city: 'Brussels',  lat: 50.8950,  lng:  4.3418  },
    { id: 10, city: 'Vienna',    lat: 48.2082,  lng: 16.3738  },
  ];

  // FIX 1: $derived value, not $derived(() => fn)
  let filtered = $derived(
    tableData.filter(r =>
      Object.values(r).some(v => String(v).toLowerCase().includes(filterText.toLowerCase()))
    )
  );

  // FIX 1 (cont): was $derived(() => { ... }) — a function, not a value
  let sorted = $derived(
    (() => {
      if (!sorting.column) return filtered;
      return [...filtered].sort((a, b) => {
        const av = a[sorting.column], bv = b[sorting.column];
        const cmp = typeof av === 'number' ? av - bv : String(av).localeCompare(String(bv));
        return sorting.direction === 'asc' ? cmp : -cmp;
      });
    })()
  );

  let totalPages = $derived(Math.max(1, Math.ceil(filtered.length / pageSize)));
  let pageRows  = $derived(sorted.slice((currentPage - 1) * pageSize, currentPage * pageSize));

  let selectedRow = $state(null);
  let markers = [];

  function setSort(col) {
    sorting = sorting.column === col
      ? { column: col, direction: sorting.direction === 'asc' ? 'desc' : 'asc' }
      : { column: col, direction: 'asc' };
  }

  function selectRow(row) {
    selectedRow = row;
    if (map) map.flyTo([row.lat, row.lng], 10, { duration: 1.2 });
  }

  onMount(async () => {
    if (typeof window === 'undefined') return;

    const leaflet = await import('https://esm.sh/leaflet@1.9.4');
    L = leaflet.default ?? leaflet;

    if (!document.getElementById('leaflet-css')) {
      const link = document.createElement('link');
      link.id = 'leaflet-css';
      link.rel = 'stylesheet';
      link.href = 'https://unpkg.com/leaflet@1.9.4/dist/leaflet.css';
      document.head.appendChild(link);
    }

    map = L.map(mapEl).setView([48.8, 8], 4);

    L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
      attribution: '© <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a>',
      maxZoom: 19,
    }).addTo(map);

    tableData.forEach(row => {
      const icon = L.divIcon({
        className: '',
        html: `<div style="width:10px;height:10px;border-radius:50%;background:var(--primary);border:2px solid var(--background);box-shadow:0 1px 6px rgba(0,0,0,.3)"></div>`,
        iconSize: [10, 10],
        iconAnchor: [5, 5],
      });
      const m = L.marker([row.lat, row.lng], { icon })
        .addTo(map)
        .bindPopup(`<strong style="color:var(--foreground)">${row.city}</strong><br><span style="color:var(--muted-foreground);font-size:11px">${row.lat.toFixed(4)}, ${row.lng.toFixed(4)}</span>`);
      markers.push({ id: row.id, marker: m });
    });
  });

  onDestroy(() => map?.remove());

  $effect(() => {
    if (map && selectedRow) {
      markers.find(m => m.id === selectedRow.id)?.marker.openPopup();
    }
  });
</script>

<!-- Full viewport, map fills everything -->
<div class="relative w-full h-full">

  <!-- MAP: fills entire container -->
  <div bind:this={mapEl} class="absolute inset-0 z-0"></div>

  <!-- ── DESKTOP: vertical pill on the right ── -->
  <div
    class="hidden md:flex absolute top-1/2 right-6 z-10 -translate-y-1/2 flex-col overflow-hidden border border-border bg-card/85 backdrop-blur-md shadow-2xl"
    style="width: 30%; height: 75vh; border-radius: 94px;"
  >
    <!-- Label + filter -->
    <div class="flex flex-shrink-0 items-center justify-between px-10 pt-8 pb-4">
      <span class="text-xs font-semibold uppercase tracking-widest text-muted-foreground">Cities</span>
      <input
        type="search"
        placeholder="Filter…"
        bind:value={filterText}
        oninput={() => { currentPage = 1; }}
        class="w-32 rounded-full border border-border bg-background/60 px-4 py-1.5 text-[12px] text-foreground placeholder:text-muted-foreground outline-none transition focus:ring-1 focus:ring-ring"
      />
    </div>

    <!-- Column headers -->
    <div class="flex-shrink-0 px-10">
      <table class="w-full border-collapse text-[13px]">
        <thead>
          <tr>
            {#each [['city','City'],['lat','Lat'],['lng','Lng']] as [col, label]}
              <th
                onclick={() => setSort(col)}
                class="cursor-pointer select-none border-b border-border pb-2.5 text-left text-[11px] font-semibold uppercase tracking-wider transition-colors
                  {sorting.column === col ? 'text-primary' : 'text-muted-foreground hover:text-foreground'}"
              >
                {label}
                <span class="ml-0.5 text-[10px] opacity-60">
                  {sorting.column === col ? (sorting.direction === 'asc' ? '↑' : '↓') : '↕'}
                </span>
              </th>
            {/each}
          </tr>
        </thead>
      </table>
    </div>

    <!-- Scrollable rows -->
    <div class="min-h-0 flex-1 overflow-y-auto px-10">
      <table class="w-full border-collapse text-[13px]">
        <tbody>
          {#each pageRows as row (row.id)}
            <tr
              onclick={() => selectRow(row)}
              class="cursor-pointer border-b border-border/50 transition-colors"
              style={selectedRow?.id === row.id ? 'background: color-mix(in oklch, var(--accent) 20%, transparent);' : ''}
              onmouseenter={e => { if (selectedRow?.id !== row.id) e.currentTarget.style.background = 'color-mix(in oklch, var(--accent) 10%, transparent)'; }}
              onmouseleave={e => { if (selectedRow?.id !== row.id) e.currentTarget.style.background = ''; }}
            >
              <td class="relative py-3 pr-3 font-medium text-foreground">
                {#if selectedRow?.id === row.id}
                  <span class="absolute inset-y-0 -left-3 w-0.5 rounded-r bg-primary"></span>
                {/if}
                {row.city}
              </td>
              <td class="py-3 pr-3 font-mono text-[12px] text-muted-foreground">{row.lat.toFixed(4)}</td>
              <td class="py-3 font-mono text-[12px] text-muted-foreground">{row.lng.toFixed(4)}</td>
            </tr>
          {/each}
          {#if pageRows.length === 0}
            <tr><td colspan="3" class="py-10 text-center text-[13px] text-muted-foreground">No results.</td></tr>
          {/if}
        </tbody>
      </table>
    </div>

    <!-- Pagination -->
    <div class="flex flex-shrink-0 items-center justify-between px-10 pt-4 pb-8">
      <span class="text-[11px] text-muted-foreground">
        {filtered.length} row{filtered.length !== 1 ? 's' : ''} · {currentPage}/{totalPages}
      </span>
      <div class="flex gap-1">
        {#each [['«', () => currentPage = 1, currentPage === 1], ['‹', () => currentPage--, currentPage === 1], ['›', () => currentPage++, currentPage === totalPages], ['»', () => currentPage = totalPages, currentPage === totalPages]] as [label, action, disabled]}
          <button onclick={action} {disabled}
            class="rounded-full border border-border bg-background/60 px-2.5 py-0.5 text-[12px] text-muted-foreground transition hover:enabled:bg-accent hover:enabled:text-accent-foreground disabled:cursor-not-allowed disabled:opacity-30">
            {label}
          </button>
        {/each}
      </div>
    </div>
  </div>

  <!-- ── MOBILE: horizontal pill along the bottom ── -->
  <!-- FIX 3: increased height from 38vh → 48vh for better data visibility -->
  <div
    class="flex md:hidden absolute bottom-6 left-1/2 z-10 -translate-x-1/2 flex-col overflow-hidden border border-border bg-card/85 backdrop-blur-md shadow-2xl"
    style="width: 92vw; height: 48vh; border-radius: 48px;"
  >
    <!-- Label + filter -->
    <div class="flex flex-shrink-0 items-center justify-between px-7 pt-5 pb-3">
      <span class="text-[10px] font-semibold uppercase tracking-widest text-muted-foreground">Cities</span>
      <input
        type="search"
        placeholder="Filter…"
        bind:value={filterText}
        oninput={() => { currentPage = 1; }}
        class="w-28 rounded-full border border-border bg-background/60 px-3 py-1 text-[11px] text-foreground placeholder:text-muted-foreground outline-none transition focus:ring-1 focus:ring-ring"
      />
    </div>

    <!-- Scrollable table -->
    <!-- FIX 2: was using `filtered` (all rows), now uses paginated `pageRows` -->
    <div class="min-h-0 flex-1 overflow-y-auto overflow-x-auto px-7">
      <table class="w-full border-collapse text-[12px]" style="min-width: max-content;">
        <thead class="sticky top-0 bg-card/95 backdrop-blur-sm">
          <tr>
            {#each [['city','City'],['lat','Lat'],['lng','Lng']] as [col, label]}
              <th
                onclick={() => setSort(col)}
                class="cursor-pointer select-none border-b border-border pb-2 pr-8 text-left text-[10px] font-semibold uppercase tracking-wider transition-colors
                  {sorting.column === col ? 'text-primary' : 'text-muted-foreground hover:text-foreground'}"
              >
                {label}
                <span class="ml-0.5 text-[9px] opacity-60">
                  {sorting.column === col ? (sorting.direction === 'asc' ? '↑' : '↓') : '↕'}
                </span>
              </th>
            {/each}
          </tr>
        </thead>
        <tbody>
          {#each pageRows as row (row.id)}
            <tr
              onclick={() => selectRow(row)}
              class="cursor-pointer border-b border-border/50 transition-colors"
              style={selectedRow?.id === row.id ? 'background: color-mix(in oklch, var(--accent) 20%, transparent);' : ''}
              onmouseenter={e => { if (selectedRow?.id !== row.id) e.currentTarget.style.background = 'color-mix(in oklch, var(--accent) 10%, transparent)'; }}
              onmouseleave={e => { if (selectedRow?.id !== row.id) e.currentTarget.style.background = ''; }}
            >
              <td class="relative py-2.5 pr-8 font-medium text-foreground whitespace-nowrap">
                {#if selectedRow?.id === row.id}
                  <span class="absolute inset-y-0 -left-2 w-0.5 rounded-r bg-primary"></span>
                {/if}
                {row.city}
              </td>
              <td class="py-2.5 pr-8 font-mono text-[11px] text-muted-foreground whitespace-nowrap">{row.lat.toFixed(4)}</td>
              <td class="py-2.5 font-mono text-[11px] text-muted-foreground whitespace-nowrap">{row.lng.toFixed(4)}</td>
            </tr>
          {/each}
          {#if pageRows.length === 0}
            <tr><td colspan="3" class="py-6 text-center text-[12px] text-muted-foreground">No results.</td></tr>
          {/if}
        </tbody>
      </table>
    </div>

    <!-- Pagination (mobile) -->
    <div class="flex flex-shrink-0 items-center justify-between px-7 pt-2 pb-5">
      <span class="text-[10px] text-muted-foreground">
        {filtered.length} cities · {currentPage}/{totalPages}
      </span>
      <div class="flex gap-1">
        {#each [['«', () => currentPage = 1, currentPage === 1], ['‹', () => currentPage--, currentPage === 1], ['›', () => currentPage++, currentPage === totalPages], ['»', () => currentPage = totalPages, currentPage === totalPages]] as [label, action, disabled]}
          <button onclick={action} {disabled}
            class="rounded-full border border-border bg-background/60 px-2.5 py-0.5 text-[11px] text-muted-foreground transition hover:enabled:bg-accent hover:enabled:text-accent-foreground disabled:cursor-not-allowed disabled:opacity-30">
            {label}
          </button>
        {/each}
      </div>
    </div>
  </div>

</div>