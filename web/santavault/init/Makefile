# Get the root directory (two levels up from src/custom)
LW_ROOT ?= $(realpath $(dir $(realpath $(lastword $(MAKEFILE_LIST))))/../../../)
LW_WORKSPACE ?= $(LW_ROOT)/workspace
LW_INSTALL ?= $(LW_WORKSPACE)/install
LW_BUILD ?= $(LW_WORKSPACE)/build

# Output directory for build artifacts
BUILD_DIR = $(LW_BUILD)/custom
INSTALL_DIR = $(LW_INSTALL)/custom

# Compiler and toolchain paths
CC = $(LW_INSTALL)/llvm/bin/clang
AR = $(LW_INSTALL)/llvm/bin/llvm-ar
LD = $(LW_INSTALL)/llvm/bin/ld.lld

# WASM target and flags (following musl and busybox patterns)
TARGET = wasm32-unknown-unknown
CFLAGS = --target=$(TARGET)
CFLAGS += -Xclang -target-feature -Xclang +atomics
CFLAGS += -Xclang -target-feature -Xclang +bulk-memory
CFLAGS += -fPIC
CFLAGS += --sysroot=$(LW_INSTALL)/musl
CFLAGS += -isystem $(LW_INSTALL)/busybox-kernel-headers
CFLAGS += -D__linux__
CFLAGS += -O2
CFLAGS += -Wall

# Linker flags (matching busybox pattern for WASM userspace)
LDFLAGS = --rtlib=compiler-rt
LDFLAGS += -nostdlib
LDFLAGS += -Wl,--export-all
LDFLAGS += -Wl,--import-table
LDFLAGS += -Wl,--import-memory
LDFLAGS += -Wl,--shared-memory
LDFLAGS += -Wl,--max-memory=4294967296
LDFLAGS += -Wl,--no-merge-data-segments
LDFLAGS += -Wl,--no-gc-sections
LDFLAGS += -Wl,--import-undefined
LDFLAGS += -Wl,-shared
LDFLAGS += -lc
# Explicitly link compiler-rt builtins for intrinsics like __multi3
LDFLAGS += $(LW_INSTALL)/llvm/lib/clang/18/lib/wasm32-unknown-unknown/libclang_rt.builtins.a

# Source files
SRCS = main.c
OBJS = $(SRCS:%.c=$(BUILD_DIR)/%.o)
TARGET_BIN = $(BUILD_DIR)/init

.PHONY: all clean install

all: $(TARGET_BIN)

$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

$(BUILD_DIR)/%.o: %.c | $(BUILD_DIR)
	$(CC) $(CFLAGS) -c $< -o $@

$(TARGET_BIN): $(OBJS)
	$(CC) $(CFLAGS) $(LDFLAGS) $(OBJS) -o $@
	@echo "Build complete: $(TARGET_BIN)"

install: $(TARGET_BIN)
	mkdir -p $(INSTALL_DIR)/sbin
	cp $(TARGET_BIN) $(INSTALL_DIR)/sbin/init
	@echo "Installed to: $(INSTALL_DIR)"

clean:
	rm -rf $(BUILD_DIR)
	@echo "Clean complete"

# Display current configuration
config:
	@echo "Configuration:"
	@echo "  LW_ROOT:      $(LW_ROOT)"
	@echo "  LW_WORKSPACE: $(LW_WORKSPACE)"
	@echo "  LW_INSTALL:   $(LW_INSTALL)"
	@echo "  LW_BUILD:     $(LW_BUILD)"
	@echo "  CC:           $(CC)"
	@echo "  TARGET:       $(TARGET)"
	@echo "  BUILD_DIR:    $(BUILD_DIR)"
	@echo "  INSTALL_DIR:  $(INSTALL_DIR)"

