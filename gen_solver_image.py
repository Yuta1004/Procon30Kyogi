import os
import glob


def make_base_image():
    check_path("Makefile")


def make_solver_image(solver_path):
    pass


def get_solver_list():
    return glob.glob("./solvers/*/solver_*.py")


def check_path(path):
    if not os.path.exists(path):
        print("   -> " + str(path) + "が存在しません")
        print("   -> プログラムを終了します")
        exit(1)


def main():
    print("Step1 : ソルバプログラムを探しています...")
    solver_list = get_solver_list()
    print("  -> " + str(len(solver_list)) + "つのソルバプログラムが見つかりました\n")

    print("Step2 : ベースイメージを生成しています...")
    make_base_image()

    print("Step3 : ソルバイメージを作成しています...")
    for solver_path in solver_list:
        make_solver_image(solver_path)


if __name__ == "__main__":
    main()
