from source_code_parser import config, server


def main():
    config.load_env_var()
    server.start()


if __name__ == "__main__":
    main()
