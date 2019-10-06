import os
import glob
import subprocess


def make_base_image():
    check_path("Makefile")
    try:
        subprocess.run(["make", "docker-build-base"])
    except subprocess.CalledProcessError:
        print("  -> ベースイメージ作成に失敗しました")
        print("  -> プログラムを終了します")
        exit(1)


def make_solver_image(solver_path):
    # Format : ./~~/~~/solver_ver1.0.py
    img_name = solver_path.split("/")[-1]                   # ファイル名
    img_name = img_name.replace(".py", "").split("_")[-1]   # バージョン
    img_name = img_name.replace(".", "")                     # .を消す
    img_name = "procon30-solver:" + img_name

    try:
        subprocess.run(["make", "docker-build-solver",
                       "SOURCE_PY="+str(solver_path), "SOLVER_IMAGE="+str(img_name)])
    except subprocess.CalledProcessError:
        print("  -> ソルバイメージ作成に失敗しました")
        print("  -> プログラムを終了します")
        exit(1)


def get_solver_list():
    check_path("solvers/")
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
    print()

    print("Step3 : ソルバイメージを作成しています...")
    for solver_path in solver_list:
        make_solver_image(solver_path)
        print()


if __name__ == "__main__":
    main()
