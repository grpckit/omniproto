# omniproto

Omniproto makes your generating gRPC code from protorepo easy.

## Running

Omniproto is bundled with [grpckit](https://github.com/grpckit/grpckit), an all-in-one docker container with a variety of plugins installed. In your grpc service repository, simply
submodule your protorepo, add an `omniproto.yaml` config file,
and a docker execute step as part of your toolchain:

`docker run -v $(pwd):/workspace --rm grpckit/omniproto`

See [omniproto.yaml](omniproto.yaml) for reference.

## How It Works

`omniproto` is designed to work with a protorepo, meaning you or your
organization has a repository dedicated to all your gRPC and protocol
buffer files. The protorepo is submoduled into your project's code
repository, and then `omniproto` is run to generate anything you need
from your protorepo. You should use `omniproto` with [grpckit](https://github.com/grpckit/grpckit)
so you have a repeatable build process across your entire environment:
CICD, developer machines, you name it.

## Configuration

`omniproto` expects a yaml or json config file in the root of your
code repository. This declares the folders or files you want to generate
from your submoduled protorepo, along with the language you want and any additional
plugins. See the `omniproto.yaml` file in the root of this repository, which
uses the protorepo-example file.

All files in your protorepo should be root-relative to the protorepo.

## A note on Google APIs

Most repositories make use of [Google APIs](https://github.com/googleapis/googleapis),
the public repository of Google, which includes `annotations` and `http`. `grpckit`
does not include these in its container, by choice. The Google APIs should be
submoduled into your protorepo, so the files you need and their location is
explicit.
