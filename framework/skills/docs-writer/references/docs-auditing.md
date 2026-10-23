# Procedural Guide: Auditing the Docset

This guide outlines the process for auditing the target project's documentation for
correctness and adherence to style guidelines. This process involves both an
"Editor" and "Technical Writer" phase.

## Objective

To ensure all public-facing documentation is accurate, up-to-date, adheres to
the target project's documentation style guide, and reflects the current state of the
codebase.

## Phase 1: Editor Audit

**Role:** The editor is responsible for identifying potential issues based on
style guide violations and technical inaccuracies.

### Steps

1.  **Identify Documentation Scope:**
    - Discover the target's actual documentation root and navigation mechanism.
      Read a sidebar or site configuration when present; otherwise inventory the
      requested Markdown/doc files directly. Resolve routes using the target's
      framework instead of assuming a slug maps to a particular filename.

2.  **Prepare Audit Results File:**
    - Create `$OUTPUT_DIR/audit-results-[YYYY-MM-DD].md`, using the main skill's
      expanded output directory and preserving existing files. It contains identified
      violations and recommendations.

3.  **Retrieve Style Guidelines:**
    - Familiarize yourself with the `docs-writer` skill instructions and the
      included style guidelines.

4.  **Audit Each Document:**
    - For each documentation file identified in Step 1, read its content.
    - **Review against Style Guide:**
      - **Voice and Tone Violations:**
        - **Unprofessional Tone:** Identify phrasing that is overly casual,
          defensive, or lacks a professional and friendly demeanor.
        - **Indirectness or Vagueness:** Identify sentences that are
          unnecessarily wordy or fail to be concise and direct.
        - **Incorrect Pronoun:** Identify any use of third-person pronouns
          (e.g., "we," "they," "the user") when referring to the reader, instead
          of the second-person pronoun **"you"**.
        - **Passive Voice:** Identify sentences written in the passive voice.
        - **Incorrect Tense:** Identify the use of past or future tense verbs,
          instead of the **present tense**.
        - **Poor Vocabulary:** Identify the use of jargon, slang, or overly
          informal language.
      - **Language and Grammar Violations:**
        - **Lack of Conciseness:** Identify unnecessarily long phrases or
          sentences.
        - **Punctuation Errors:** Identify incorrect or missing punctuation.
        - **Ambiguous Dates:** Identify dates that could be misinterpreted
          (e.g., "next Monday" instead of "April 15, 2026").
        - **Abbreviation Usage:** Identify the use of abbreviations that should
          be spelled out (e.g., "e.g." instead of "for example").
        - **Terminology:** Check for incorrect or inconsistent use of
          product-specific terms (e.g., "quota" vs. "limit").
      - **Formatting and Syntax Violations:**
        - **Missing Overview:** Check for the absence of a brief overview
          paragraph at the start of the document.
        - **Line Length:** Identify any lines of text that exceed **80
          characters** (text wrap violation).
        - **Casing:** Identify incorrect casing for headings, titles, or named
           entities (use the target's documented product name).
        - **List Formatting:** Identify incorrectly formatted lists (e.g.,
          inconsistent indentation or numbering).
        - **Incorrect Emphasis:** Identify incorrect use of bold text (should
          only be used for UI elements) or code font (should be used for code,
          file names, or command-line input).
        - **Link Quality:** Identify links with non-descriptive anchor text
          (e.g., "click here").
        - **Image Alt Text:** Identify images with missing or poor-quality
          (non-descriptive) alt text.
      - **Structure Violations:**
        - **Missing BLUF:** Check for the absence of a "Bottom Line Up Front"
          summary at the start of complex sections or documents.
        - **Experimental Feature Notes:** Identify experimental features that
          are not clearly labeled with a standard note.
        - **Heading Hierarchy:** Check for skipped heading levels (e.g., going
          from `##` to `####`).
        - **Procedure Clarity:** Check for procedural steps that do not start
          with an imperative verb or where a condition is placed _after_ the
          instruction.
        - **Element Misuse:** Identify the incorrect or inappropriate use of
          special elements (e.g., Notes, Warnings, Cautions).
        - **Table of Contents:** Identify the presence of a dynamically
          generated or manually included table of contents.
        - **Missing Next Steps:** Check for procedural documents that lack a
          "Next steps" section (if applicable).
    - **Verify Code Accuracy (if applicable):**
      - If the document contains code snippets (e.g., shell commands, API calls,
       file paths, Docker image versions), use the host's search and read tools
         within the relevant parts of the
        codebase) to ensure the code is still accurate and up-to-date. Pay close
        attention to version numbers, package names, and command syntax.
    - **Record Findings:** For each **violation** or inaccuracy found:
      - Note the file path.
      - Describe the violation (e.g., "Violation (Language and Grammar): Uses
        'e.g.'").
      - Provide a clear and actionable recommendation to correct the issue.
        (e.g., "Recommendation: Replace 'e.g.' with 'for example'." or
        "Recommendation: Replace '...' with '...' in active voice.).
      - Append these findings to `audit-results-[YYYY-MM-DD].md`.

## Phase 2: Software Engineer Audit

**Role:** The software engineer is responsible for finding undocumented features
by auditing the codebase and recent changelogs, and passing these findings to
the technical writer.

### Steps

1.  **Proactive Codebase Audit:**
     - Audit the target's entry points, public APIs, configuration, commands,
       and user-visible workflows. Discover their actual locations from the
       repository rather than assuming Gemini CLI's package layout.

2.  **Review Recent Updates:**
    - Check recent changelogs in stable and announcements within the
      documentation to see if newly introduced features are documented properly.

3.  **Evaluate and Record Findings:**
    - Determine if these features are adequately covered in the docs. They do
      not need to be documented word for word, but major features that customers
      should care about probably should have an article.
    - Append your findings to the `audit-results-[YYYY-MM-DD].md` file,
      providing a brief description of the feature and where it should be
      documented.

## Phase 3: Technical Writer Implementation

**Role:** The technical writer handles input from both the editor and the
software engineer, makes appropriate decisions about what to change, and
implements the approved changes.

### Steps

1.  **Review Audit Results:**
     - Read `$OUTPUT_DIR/audit-results-[YYYY-MM-DD].md` to understand all identified issues,
      undocumented features, and recommendations from both the Editor and
      Software Engineer phases.

2.  **Make Decisions and Log Reasoning:**
     - Create or update `$OUTPUT_DIR/audit-implementation-log-[YYYY-MM-DD].md`
       only when implementing changes, preserving existing reports.
    - Make sure the logs are updated for all steps, documenting your reasoning
      for each recommendation (why it was accepted, modified, or rejected). This
      is required for a final check by a human in the PR.

3.  **Implement Changes:**
    - For each approved recommendation:
      - Read the target documentation file.
       - Apply scoped changes with the current host's editing tool. Match the
         intended context exactly and review every replacement; do not use blind
         whole-tree string substitutions.
      - **String replacement safeguards:** When applying these fixes across the
        docset, you must verify the following:
        - **Preserve Code Blocks:** Explicitly verify that no code blocks,
          inline code snippets, terminal commands, or file paths have been
          erroneously capitalized or modified.
        - **Preserve Literal Strings:** Never alter the wording of literal error
          messages, UI quotes, or system logs. For example, if a style rule says
          to remove the word "please", you must NOT remove it if it appears
          inside a quoted error message (e.g.,
          `Error: Please contact your administrator`).
        - **Verify Sentence Casing:** When removing filler words (like "please")
          from the beginning of a sentence or list item, always verify that the
          new first word of the sentence is properly capitalized.
       - For structural changes, use the host's normal file-editing tools.
      - For broken links, determine the correct new path or update the link
        text.
      - For creating new files (e.g., `docs/get-started.md` to fix a broken
         link, or a new feature article), follow the project's contribution process.

4.  **Execute Auto-Generation Scripts:**
     - Discover the target's documented generators and use those when relevant
       and authorized. Do not assume npm, settings, or keybinding generators exist.

5.  **Format Code:**
     - Run the target's own documentation and formatting checks when available.
       Do not install dependencies or invent a formatter without approval. State
       what ran and any blocked checks. An audit-only request changes no docs.
