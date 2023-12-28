# Generated by the protocol buffer compiler.  DO NOT EDIT!
# sources: queue.proto
# plugin: python-betterproto
from dataclasses import dataclass

import betterproto
import grpclib


@dataclass
class SimpleQueueEnqueueRequest(betterproto.Message):
    name: str = betterproto.string_field(1)
    value: bytes = betterproto.bytes_field(2)


@dataclass
class SimpleQueueEnqueueResponse(betterproto.Message):
    ok: bool = betterproto.bool_field(1)


@dataclass
class SimpleQueueDequeueRequest(betterproto.Message):
    name: str = betterproto.string_field(1)
    value: bytes = betterproto.bytes_field(2)


@dataclass
class SimpleQueueDequeueResponse(betterproto.Message):
    ok: bool = betterproto.bool_field(1)
    value: bytes = betterproto.bytes_field(2)


@dataclass
class SimpleQueuePeekResponse(betterproto.Message):
    ok: bool = betterproto.bool_field(1)
    value: bytes = betterproto.bytes_field(2)


@dataclass
class SimpleQueueEmptyResponse(betterproto.Message):
    ok: bool = betterproto.bool_field(1)
    empty: bool = betterproto.bool_field(2)


@dataclass
class SimpleQueueSizeResponse(betterproto.Message):
    ok: bool = betterproto.bool_field(1)
    size: int = betterproto.uint64_field(2)


@dataclass
class SimpleQueueRequest(betterproto.Message):
    name: str = betterproto.string_field(1)


class SimpleQueueServiceStub(betterproto.ServiceStub):
    async def enqueue(
        self, *, name: str = "", value: bytes = b""
    ) -> SimpleQueueEnqueueResponse:
        request = SimpleQueueEnqueueRequest()
        request.name = name
        request.value = value

        return await self._unary_unary(
            "/simplequeue.SimpleQueueService/Enqueue",
            request,
            SimpleQueueEnqueueResponse,
        )

    async def dequeue(
        self, *, name: str = "", value: bytes = b""
    ) -> SimpleQueueDequeueResponse:
        request = SimpleQueueDequeueRequest()
        request.name = name
        request.value = value

        return await self._unary_unary(
            "/simplequeue.SimpleQueueService/Dequeue",
            request,
            SimpleQueueDequeueResponse,
        )

    async def peek(self, *, name: str = "") -> SimpleQueuePeekResponse:
        request = SimpleQueueRequest()
        request.name = name

        return await self._unary_unary(
            "/simplequeue.SimpleQueueService/Peek",
            request,
            SimpleQueuePeekResponse,
        )

    async def empty(self, *, name: str = "") -> SimpleQueueEmptyResponse:
        request = SimpleQueueRequest()
        request.name = name

        return await self._unary_unary(
            "/simplequeue.SimpleQueueService/Empty",
            request,
            SimpleQueueEmptyResponse,
        )

    async def size(self, *, name: str = "") -> SimpleQueueSizeResponse:
        request = SimpleQueueRequest()
        request.name = name

        return await self._unary_unary(
            "/simplequeue.SimpleQueueService/Size",
            request,
            SimpleQueueSizeResponse,
        )