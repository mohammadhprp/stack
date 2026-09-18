package harness

import (
	"io/fs"
	"path/filepath"

	"github.com/mohammadhprp/stack/internal/config"
	"github.com/mohammadhprp/stack/internal/models"
)

func reconcileJSONMCPs(target, configFile, section string, keep, drop []models.MCP, render func(models.MCP) (map[string]any, error)) ([]File, []string, error) {
	entries := make(map[string]any, len(keep))
	for _, mcp := range keep {
		entry, err := render(mcp)
		if err != nil {
			return nil, nil, err
		}
		entries[mcp.ID] = entry
	}
	dropIDs := make([]string, 0, len(drop))
	for _, mcp := range drop {
		dropIDs = append(dropIDs, mcp.ID)
	}
	out, remove, err := config.ReconcileJSONSection(filepath.Join(target, filepath.FromSlash(configFile)), section, entries, dropIDs)
	if err != nil {
		return nil, nil, err
	}
	if remove {
		return nil, []string{configFile}, nil
	}
	return []File{{Path: configFile, Content: out, Merge: true, Source: "mcps"}}, nil, nil
}

func reconcileTOMLMCPs(target, configFile, table string, keep, drop []models.MCP, render func(models.MCP) (map[string]any, error)) ([]File, []string, error) {
	entries := make(map[string]any, len(keep))
	for _, mcp := range keep {
		entry, err := render(mcp)
		if err != nil {
			return nil, nil, err
		}
		entries[mcp.ID] = entry
	}
	dropIDs := make([]string, 0, len(drop))
	for _, mcp := range drop {
		dropIDs = append(dropIDs, mcp.ID)
	}
	out, remove, err := config.ReconcileTOMLTable(filepath.Join(target, filepath.FromSlash(configFile)), table, entries, dropIDs)
	if err != nil {
		return nil, nil, err
	}
	if remove {
		return nil, []string{configFile}, nil
	}
	return []File{{Path: configFile, Content: out, Merge: true, Source: "mcps"}}, nil, nil
}

func (opencode) PlanMCPRemoval(target string, keep, drop []models.MCP, _ fs.FS) ([]File, []string, error) {
	return reconcileJSONMCPs(target, opencodeConfig, "mcp", keep, drop, renderOpenCodeMCP)
}

func (claude) PlanMCPRemoval(target string, keep, drop []models.MCP, _ fs.FS) ([]File, []string, error) {
	return reconcileJSONMCPs(target, claudeConfig, "mcpServers", keep, drop, claudeMCPEntry)
}

func (cursor) PlanMCPRemoval(target string, keep, drop []models.MCP, _ fs.FS) ([]File, []string, error) {
	return reconcileJSONMCPs(target, cursorConfig, "mcpServers", keep, drop, cursorMCPEntry)
}

func (gemini) PlanMCPRemoval(target string, keep, drop []models.MCP, _ fs.FS) ([]File, []string, error) {
	return reconcileJSONMCPs(target, geminiConfig, "mcpServers", keep, drop, geminiMCPEntry)
}

func (amp) PlanMCPRemoval(target string, keep, drop []models.MCP, _ fs.FS) ([]File, []string, error) {
	return reconcileJSONMCPs(target, ampConfig, "amp.mcpServers", keep, drop, ampMCPEntry)
}

func (codex) PlanMCPRemoval(target string, keep, drop []models.MCP, _ fs.FS) ([]File, []string, error) {
	return reconcileTOMLMCPs(target, codexConfig, "mcp_servers", keep, drop, codexMCPEntry)
}
