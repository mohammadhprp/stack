# Figma MCP Go Capabilities

## What It Can Do

73 tools total across read, write, export, and prompt categories.

### Write — Create

| Tool | Description |
| --- | --- |
| `create_frame` | Create a frame with optional auto-layout, fill, and parent |
| `create_rectangle` | Create a rectangle with optional fill and corner radius |
| `create_ellipse` | Create an ellipse or circle |
| `create_text` | Create a text node (font loaded automatically) |
| `import_image` | Decode base64 image and place it as a rectangle fill |
| `create_component` | Convert an existing FRAME node into a reusable component |
| `create_section` | Create a Figma Section node to organise frames on a page |

### Write — Modify

| Tool | Description |
| --- | --- |
| `set_text` | Update text content of an existing TEXT node |
| `set_fills` | Set solid fill color (hex) on a node |
| `set_strokes` | Set solid stroke color and weight on a node |
| `set_opacity` | Set opacity of one or more nodes |
| `set_corner_radius` | Set corner radius — uniform or per-corner |
| `set_auto_layout` | Set or update auto-layout (flex) properties on a frame |
| `set_visible` | Show or hide one or more nodes |
| `lock_nodes` | Lock one or more nodes to prevent accidental edits |
| `unlock_nodes` | Unlock one or more nodes |
| `rotate_nodes` | Set absolute rotation in degrees on one or more nodes |
| `reorder_nodes` | Change z-order |
| `set_blend_mode` | Set blend mode on one or more nodes |
| `set_constraints` | Set responsive constraints on one or more nodes |
| `move_nodes` | Move nodes to an absolute x/y position |
| `resize_nodes` | Resize nodes by width and/or height |
| `rename_node` | Rename a node |
| `clone_node` | Clone a node, optionally repositioning or reparenting |
| `reparent_nodes` | Move nodes to a different parent |
| `batch_rename_nodes` | Bulk rename nodes via find/replace, regex, or prefix/suffix |
| `find_replace_text` | Find and replace text across all TEXT nodes |

### Write — Delete

| Tool | Description |
| --- | --- |
| `delete_nodes` | Delete one or more nodes permanently |

### Write — Prototype

| Tool | Description |
| --- | --- |
| `set_reactions` | Set prototype reactions (triggers + actions) on a node |
| `remove_reactions` | Remove all or specific reactions from a node |

### Write — Styles

| Tool | Description |
| --- | --- |
| `set_effects` | Apply drop shadow / blur effects directly on a node |
| `create_paint_style` | Create a named paint style with a solid color |
| `create_text_style` | Create a named text style |
| `create_effect_style` | Create a named effect style |
| `create_grid_style` | Create a named layout grid style |
| `update_paint_style` | Rename or recolor an existing paint style |
| `apply_style_to_node` | Apply an existing local style to a node |
| `delete_style` | Delete any style by ID |

### Write — Variables

| Tool | Description |
| --- | --- |
| `create_variable_collection` | Create a new local variable collection |
| `add_variable_mode` | Add a new mode to an existing collection |
| `create_variable` | Create a variable in a collection |
| `set_variable_value` | Set a variable's value for a specific mode |
| `bind_variable_to_node` | Bind a variable to a node property |
| `delete_variable` | Delete a variable or an entire collection |

### Write — Pages

| Tool | Description |
| --- | --- |
| `add_page` | Add a new page to the document |
| `delete_page` | Delete a page by ID or name |
| `rename_page` | Rename a page by ID or current name |

### Write — Components & Navigation

| Tool | Description |
| --- | --- |
| `navigate_to_page` | Switch the active Figma page |
| `group_nodes` | Group two or more nodes into a GROUP |
| `ungroup_nodes` | Ungroup GROUP nodes |
| `swap_component` | Swap the main component of an INSTANCE node |
| `detach_instance` | Detach component instances to plain frames |

### Read — Document & Selection

| Tool | Description |
| --- | --- |
| `get_document` | Full current page tree |
| `get_metadata` | File name, pages, current page |
| `get_pages` | All pages (IDs + names) — lightweight |
| `get_selection` | Currently selected nodes |
| `get_node` | Single node by ID |
| `get_nodes_info` | Multiple nodes by ID |
| `get_design_context` | Depth-limited tree with detail level |
| `search_nodes` | Find nodes by name and/or type |
| `scan_text_nodes` | All text nodes in a subtree |
| `scan_nodes_by_types` | Nodes matching given type list |
| `get_viewport` | Current viewport center, zoom, and visible bounds |

### Read — Styles & Variables

| Tool | Description |
| --- | --- |
| `get_styles` | Paint, text, effect, and grid styles |
| `get_variable_defs` | Variable collections and values |
| `get_local_components` | All components + component sets |
| `get_annotations` | Dev-mode annotations |
| `get_fonts` | All fonts used on the current page |
| `get_reactions` | Prototype/interaction reactions on a node |

### Export

| Tool | Description |
| --- | --- |
| `get_screenshot` | Base64 image export of any node |
| `save_screenshots` | Export images to disk |
| `export_frames_to_pdf` | Export multiple frames as a single multi-page PDF |
| `export_tokens` | Export design tokens as JSON or CSS |

### MCP Prompts

| Prompt | Description |
| --- | --- |
| `read_design_strategy` | Best practices for reading Figma designs |
| `design_strategy` | Best practices for creating and modifying designs |
| `text_replacement_strategy` | Chunked approach for replacing text across a design |
| `annotation_conversion_strategy` | Convert manual annotations to native Figma annotations |
| `swap_overrides_instances` | Transfer overrides between component instances |
| `reaction_to_connector_strategy` | Map prototype reactions into interaction flow diagrams |

## What It Cannot Do

- It cannot work without the Figma desktop plugin running inside an open file.
- It does not use the Figma REST API, so it cannot read files or projects without the plugin bridge.
- It cannot access Figma team libraries, shared styles, or remote variables outside the current file.
- It cannot replace human review for destructive operations like deleting nodes or bulk renaming.
- It cannot export SVGs or vector files directly; exports are limited to images and PDFs.

## Best Practices

- Keep the Figma plugin running in the file being edited.
- Prefer `get_document` or `get_design_context` to understand the structure before making changes.
- Use `search_nodes` to find specific elements by name rather than traversing the full tree.
- Use `batch_rename_nodes` and `find_replace_text` for bulk changes instead of per-node operations.
- Verify destructive operations (`delete_nodes`, `delete_style`, `delete_variable`) before executing.
- Run exports (`save_screenshots`, `export_frames_to_pdf`) to disk for durable artifacts.
- Keep the plugin version in sync with the MCP server version.

## Common Workflows

### Understand a Figma design

1. Ask the agent to call `get_metadata` to learn file and page structure.
2. Call `get_document` or `get_design_context` on the relevant page.
3. Call `get_styles`, `get_variable_defs`, and `get_local_components` for design system details.
4. Ask the agent to use the `read_design_strategy` prompt for guidance.

### Convert a design to code

1. Select the target frame in Figma.
2. Ask the agent to call `get_selection` and `get_node` on the selected frame.
3. Export the frame with `get_screenshot` for visual reference.
4. Use the structure and styles to generate matching code.

### Update text across a design

1. Use `find_replace_text` or `scan_text_nodes` to locate all text nodes.
2. Use `set_text` to update individual nodes or `batch_rename_nodes` for bulk operations.
3. Ask the agent to use the `text_replacement_strategy` prompt for large text replacement tasks.

### Export design tokens

1. Call `get_variable_defs` to retrieve variable collections and values.
2. Call `get_styles` to retrieve paint, text, and effect styles.
3. Use `export_tokens` to generate JSON or CSS output.